// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

// Package app implements a server that runs a set of active components.
package app

import (
	"context"
	"fmt"
	"os"

	"github.com/jinzhu/copier"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	genericapiserver "k8s.io/apiserver/pkg/server"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	restclient "k8s.io/client-go/rest"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/cli/globalflag"
	"k8s.io/component-base/logs"
	logsapi "k8s.io/component-base/logs/api/v1"
	"k8s.io/component-base/metrics/features"
	"k8s.io/component-base/term"
	"k8s.io/klog/v2"
	"k8s.io/utils/ptr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"

	"github.com/onexstack/onex/cmd/onex-blockchain-controller/app/config"
	"github.com/onexstack/onex/cmd/onex-blockchain-controller/app/options"
	chaincontroller "github.com/onexstack/onex/internal/controller/blockchain/chain"
	minercontroller "github.com/onexstack/onex/internal/controller/blockchain/miner"
	minersetcontroller "github.com/onexstack/onex/internal/controller/blockchain/minerset"
	synccontroller "github.com/onexstack/onex/internal/controller/blockchain/sync"
	resourcecleancontroller "github.com/onexstack/onex/internal/controller/resourceclean"
	"github.com/onexstack/onex/internal/gateway/store"
	"github.com/onexstack/onex/internal/pkg/util/ratelimiter"
	"github.com/onexstack/onex/internal/webhooks"
	"github.com/onexstack/onex/pkg/apis/apps/v1beta1"
	"github.com/onexstack/onex/pkg/apis/apps/v1beta1/index"
	"github.com/onexstack/onex/pkg/record"
	"github.com/onexstack/onexstack/pkg/db"
	"github.com/onexstack/onexstack/pkg/version"
)

func init() {
	utilruntime.Must(logsapi.AddFeatureGates(utilfeature.DefaultMutableFeatureGate))
	utilruntime.Must(features.AddFeatureGates(utilfeature.DefaultMutableFeatureGate))
}

// NewControllerCommand creates a *cobra.Command object with default parameters.
func NewControllerCommand() *cobra.Command {
	o, err := options.NewOptions()
	if err != nil {
		klog.Background().Error(err, "Unable to initialize command options")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}

	cmd := &cobra.Command{
		Use: "onex-blockchain-controller",
		Long: `The cloud miner controller is a daemon that embeds
the core control loops. In applications of robotics and
automation, a control loop is a non-terminating loop that regulates the state of
the system. In OneX, a controller is a control loop that watches the shared
state of the miner through the onex-apiserver and makes changes attempting to move the
current state towards the desired state.`,
		PersistentPreRunE: func(*cobra.Command, []string) error {
			// silence client-go warnings.
			// onex-blockchain-controller generically watches APIs (including deprecated ones),
			// and CI ensures it works properly against matching onex-apiserver versions.
			restclient.SetDefaultWarningHandler(restclient.NoWarnings{})
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			version.PrintAndExitIfRequested()

			// Activate logging as soon as possible, after that
			// show flags with the final logging configuration.
			if err := logsapi.ValidateAndApply(o.Logs, utilfeature.DefaultFeatureGate); err != nil {
				return err
			}
			ctrl.SetLogger(klog.Background())

			cliflag.PrintFlags(cmd.Flags())

			// klog.Background will automatically use the right logger. Here use the
			// global klog.logging initialized by `logsapi.ValidateAndApply`.
			if err := o.Complete(); err != nil {
				return err
			}

			if err := o.Validate(); err != nil {
				return err
			}

			c, err := o.Config()
			if err != nil {
				return err
			}

			if err := options.LogOrWriteConfig(o.WriteConfigTo, c.ComponentConfig); err != nil {
				return err
			}

			// add feature enablement metrics
			utilfeature.DefaultMutableFeatureGate.AddMetrics()
			return Run(genericapiserver.SetupSignalContext(), c)
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}

	fs := cmd.Flags()
	namedFlagSets := o.Flags()
	version.AddFlags(namedFlagSets.FlagSet("global"))
	globalflag.AddGlobalFlags(namedFlagSets.FlagSet("global"), cmd.Name(), logs.SkipLoggingConfigurationFlags())
	for _, f := range namedFlagSets.FlagSets {
		fs.AddFlagSet(f)
	}

	cols, _, _ := term.TerminalSize(cmd.OutOrStdout())
	cliflag.SetUsageAndHelpFunc(cmd, namedFlagSets, cols)

	if err := cmd.MarkFlagFilename("config", "yaml", "yml", "json"); err != nil {
		klog.Background().Error(err, "Failed to mark flag filename")
	}

	return cmd
}

// Run runs the controller options. This should never exit.
func Run(ctx context.Context, c *config.Config) error {
	// To help debugging, immediately log version
	klog.InfoS("Starting miner controller", "version", version.Get().String())

	klog.InfoS("Golang settings", "GOGC", os.Getenv("GOGC"), "GOMAXPROCS", os.Getenv("GOMAXPROCS"), "GOTRACEBACK", os.Getenv("GOTRACEBACK"))

	// Do some initialization here
	var redisOptions db.RedisOptions
	_ = copier.Copy(&redisOptions, &c.ComponentConfig.Redis)
	rdb, err := db.NewRedis(&redisOptions)
	if err != nil {
		return err
	}

	var mysqlOptions db.MySQLOptions
	_ = copier.Copy(&mysqlOptions, c.ComponentConfig.MySQL)
	storeClient, err := wireStoreClient(&mysqlOptions)
	if err != nil {
		return err
	}

	var watchNamespaces map[string]cache.Config
	if c.ComponentConfig.Generic.Namespace != "" {
		watchNamespaces = map[string]cache.Config{
			c.ComponentConfig.Generic.Namespace: {},
		}
	}

	ctrlOptions := ctrl.Options{
		LeaderElection:             c.ComponentConfig.Generic.LeaderElection.LeaderElect,
		LeaderElectionID:           c.ComponentConfig.Generic.LeaderElection.ResourceName,
		LeaseDuration:              &c.ComponentConfig.Generic.LeaderElection.LeaseDuration.Duration,
		RenewDeadline:              &c.ComponentConfig.Generic.LeaderElection.RenewDeadline.Duration,
		RetryPeriod:                &c.ComponentConfig.Generic.LeaderElection.RetryPeriod.Duration,
		LeaderElectionResourceLock: c.ComponentConfig.Generic.LeaderElection.ResourceLock,
		LeaderElectionNamespace:    c.ComponentConfig.Generic.LeaderElection.ResourceNamespace,
		HealthProbeBindAddress:     c.ComponentConfig.Generic.HealthzBindAddress,
		Metrics: metricsserver.Options{
			SecureServing: false,
			BindAddress:   c.ComponentConfig.Generic.MetricsBindAddress,
		},
		Cache: cache.Options{
			DefaultNamespaces: watchNamespaces,
			SyncPeriod:        &c.ComponentConfig.Generic.SyncPeriod.Duration,
		},
	}

	cctx, err := CreateControllerContext(ctx, c, storeClient, rdb, ctrlOptions)
	if err != nil {
		return err
	}

	// Create a new Cmd to provide shared dependencies and start components
	mgr, err := ctrl.NewManager(c.Kubeconfig, ctrlOptions)
	if err != nil {
		klog.ErrorS(err, "Unable to new blockchain controller")
		return err
	}

	// applies all the stored functions to the scheme created by controller-runtime
	_ = v1beta1.AddToScheme(mgr.GetScheme())
	_ = corev1.AddToScheme(mgr.GetScheme())

	// Initialize event recorder.
	record.InitFromRecorder(mgr.GetEventRecorderFor("onex-blockchain-controller"))

	if err := index.AddDefaultIndexes(ctx, mgr); err != nil {
		klog.ErrorS(err, "Unable to setup indexes")
		return err
	}

	if !c.ComponentConfig.DryRun {
		// controller-runtime for multi-cluster support, reference:
		// https://github.com/kubernetes-sigs/controller-runtime/blob/main/designs/move-cluster-specific-code-out-of-manager.md
		if err := mgr.Add(c.ProviderCluster); err != nil {
			return err
		}
	}

	setupChecks(mgr)

	if err := setupReconcilers(ctx, mgr, cctx); err != nil {
		return err
	}
	// setupWebhooks(mgr)

	return mgr.Start(ctx)
}

func setupChecks(mgr ctrl.Manager) {
	if err := mgr.AddReadyzCheck("healthz", healthz.Ping); err != nil {
		klog.Exitf("Unable to set up health check: %v", err)
	}

	if err := mgr.AddHealthzCheck("readyz", healthz.Ping); err != nil {
		klog.Exitf("Unable to set up ready check: %v", err)
	}

	/*
		// NOTICE: If you have not set up a webhook, there is no need to enable the webhook health check endpoint.
		if err := mgr.AddReadyzCheck("webhook", mgr.GetWebhookServer().StartedChecker()); err != nil {
			klog.Errorf("Unable to set up webhook ready check: %v", err)
		}

		if err := mgr.AddHealthzCheck("webhook", mgr.GetWebhookServer().StartedChecker()); err != nil {
			klog.Errorf("Unable to set up webhook health check: %v", err)
		}
	*/
}

//nolint:unused
func setupWebhooks(mgr ctrl.Manager) {
	if err := (&webhooks.Chain{}).SetupWebhookWithManager(mgr); err != nil {
		klog.Exitf("Unable to create Chain webhook: %v", err)
	}
}

func setupReconcilers(ctx context.Context, mgr ctrl.Manager, cctx ControllerContext) error {
	// Setup chain controller
	if err := (&chaincontroller.Reconciler{
		ComponentConfig:  &cctx.Config.ComponentConfig.ChainController,
		WatchFilterValue: cctx.Config.ComponentConfig.Generic.WatchFilterValue,
	}).SetupWithManager(ctx, mgr, cctx.ControllerOptions); err != nil {
		klog.ErrorS(err, "Unable to create controller", "controller", "chain")
		return err
	}
	// Setup chain sync controller
	if err := (&synccontroller.ChainSyncReconciler{
		Store: cctx.Store,
	}).SetupWithManager(ctx, mgr, cctx.ControllerOptions); err != nil {
		klog.ErrorS(err, "Unable to create controller", "controller", "chain-sync")
		return err
	}

	// Setup minerset controller
	if err := (&minersetcontroller.Reconciler{
		WatchFilterValue: cctx.Config.ComponentConfig.Generic.WatchFilterValue,
	}).SetupWithManager(ctx, mgr, controller.Options{
		MaxConcurrentReconciles: int(cctx.Config.ComponentConfig.Generic.Parallelism),
		RecoverPanic:            ptr.To(true),
		RateLimiter:             ratelimiter.DefaultControllerRateLimiter(),
	}); err != nil {
		klog.ErrorS(err, "Unable to create controller", "controller", "minerset")
		return err
	}

	// Setup minerset controller
	if err := (&synccontroller.MinerSetSyncReconciler{
		Store: cctx.Store,
	}).SetupWithManager(ctx, mgr, cctx.ControllerOptions); err != nil {
		klog.ErrorS(err, "Unable to create controller", "controller", "minerset-sync")
		return err
	}

	// Setup miner controller
	if err := (&minercontroller.Reconciler{
		DryRun:           cctx.Config.ComponentConfig.DryRun,
		ProviderClient:   cctx.Config.ProviderClient,
		ProviderCluster:  cctx.Config.ProviderCluster,
		RedisClient:      cctx.RedisClient,
		ComponentConfig:  cctx.Config.ComponentConfig,
		WatchFilterValue: cctx.Config.ComponentConfig.Generic.WatchFilterValue,
	}).SetupWithManager(ctx, mgr, controller.Options{
		MaxConcurrentReconciles: int(cctx.Config.ComponentConfig.Generic.Parallelism),
		RecoverPanic:            ptr.To(true),
		RateLimiter:             ratelimiter.DefaultControllerRateLimiter(),
	}); err != nil {
		klog.ErrorS(err, "Unable to create controller", "controller", "miner")
		return err
	}

	// Setup miner controller
	if err := (&synccontroller.MinerSyncReconciler{
		Store: cctx.Store,
	}).SetupWithManager(ctx, mgr, cctx.ControllerOptions); err != nil {
		klog.ErrorS(err, "Unable to create controller", "controller", "miner-sync")
		return err
	}

	// Setup resource clean controller
	mgr.Add(resourcecleancontroller.NewCleanReconciler(
		mgr.GetClient(),
		cctx.Store,
		&resourcecleancontroller.Miner{},
		&resourcecleancontroller.MinerSet{},
		&resourcecleancontroller.Chain{},
	))

	return nil
}

// ControllerContext defines the context object for controller
type ControllerContext struct {
	// Config provides access to init options for a given controller
	Config *config.Config

	RedisClient *redis.Client

	ControllerOptions controller.Options

	Store store.IStore
}

func CreateControllerContext(
	ctx context.Context,
	c *config.Config,
	storeClient store.IStore,
	rdb *redis.Client,
	ctrlOptions ctrl.Options,
) (ControllerContext, error) {
	return ControllerContext{
		Config:      c,
		RedisClient: rdb,
		ControllerOptions: controller.Options{
			MaxConcurrentReconciles: int(c.ComponentConfig.Generic.Parallelism),
			RecoverPanic:            ptr.To(true),
			RateLimiter:             ratelimiter.DefaultControllerRateLimiter(),
		},
		Store: storeClient,
	}, nil
}
