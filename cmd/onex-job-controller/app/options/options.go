// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

// Package options provides the flags used for the job controller.
package options

import (
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	cliflag "k8s.io/component-base/cli/flag"
	componentbaseoptions "k8s.io/component-base/config/options"
	"k8s.io/component-base/logs"
	logsapi "k8s.io/component-base/logs/api/v1"

	controllerconfig "github.com/onexstack/onex/cmd/onex-job-controller/app/config"
	jobcontrollerconfig "github.com/onexstack/onex/internal/controller/job/apis/config"
	"github.com/onexstack/onex/internal/controller/job/apis/config/latest"
	"github.com/onexstack/onex/internal/controller/job/apis/config/validation"
	clientcmdutil "github.com/onexstack/onex/internal/pkg/util/clientcmd"
	kubeutil "github.com/onexstack/onex/internal/pkg/util/kube"
)

const (
	// ControllerUserAgent is the userAgent name when starting onex-job controller.
	ControllerUserAgent = "onex-job-controller"
)

// Options is the main context object for the onex-job-controller.
type Options struct {
	// ConfigFile is the location of the job controller server's configuration file.
	ConfigFile string

	// WriteConfigTo is the path where the default configuration will be written.
	WriteConfigTo string

	// The address of the Kubernetes API server (overrides any value in kubeconfig).
	Master string

	// Path to kubeconfig file with authorization and master location information.
	Kubeconfig string

	Logs *logs.Options

	// config is the job controller server's configuration object.
	// The default values.
	config *jobcontrollerconfig.JobControllerConfiguration
}

// NewOptions creates a new Options with a default config.
func NewOptions() (*Options, error) {
	o := Options{
		Kubeconfig: clientcmdutil.DefaultKubeconfig(),
		Logs:       logs.NewOptions(),
	}

	defaultComponentConfig, err := latest.Default()
	if err != nil {
		return nil, err
	}
	o.config = defaultComponentConfig

	return &o, nil
}

// Complete completes all the required options.
func (o *Options) Complete() error {
	if len(o.ConfigFile) == 0 {
		// If the --config arg is not specified, honor the deprecated as well as leader election CLI args.
		o.ApplyDeprecated()
	} else {
		cfg, err := LoadConfigFromFile(o.ConfigFile)
		if err != nil {
			return err
		}
		o.config = cfg
	}

	return nil
}

func (o *Options) ApplyDeprecated() {}

// Flags returns flags for a specific APIServer by section name.
func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	// o.Logs.AddFlags(fss.FlagSet("logs"))
	componentbaseoptions.BindLeaderElectionFlags(&o.config.Generic.LeaderElection, fss.FlagSet("leader elect"))
	///o.config.Cloud.AddFlags(fss.FlagSet("cloud"))

	fs := fss.FlagSet("misc")
	fs.StringVar(&o.ConfigFile, "config", o.ConfigFile, "The path to the configuration file.")
	fs.StringVar(&o.WriteConfigTo, "write-config-to", o.WriteConfigTo, "If set, write the default configuration values to this file and exit.")
	fs.StringVar(&o.Master, "master", o.Master, "The address of the Kubernetes API server (overrides any value in kubeconfig).")
	fs.StringVar(&o.Kubeconfig, "kubeconfig", o.Kubeconfig, "Path to kubeconfig file with authorization and master location information.")

	logsapi.AddFlags(o.Logs, fss.FlagSet("logs"))
	utilfeature.DefaultMutableFeatureGate.AddFlag(fss.FlagSet("generic"))

	return fss
}

// Validate is used to validate the options and config before launching the controller.
func (o *Options) Validate() error {
	var errs []error

	if err := validation.Validate(o.config).ToAggregate(); err != nil {
		errs = append(errs, err.Errors()...)
	}
	// errs = append(errs, o.config.ChainController.Validate()...)

	// errs = append(errs, o.Cloud.Validate()...)

	return utilerrors.NewAggregate(errs)
}

// ApplyTo fills up job controller config with options.
func (o *Options) ApplyTo(c *controllerconfig.Config) error {
	c.ComponentConfig = o.config
	return nil
}

// Config return a job controller config objective.
func (o *Options) Config() (*controllerconfig.Config, error) {
	kubeconfig, err := clientcmd.BuildConfigFromFlags(o.Master, o.Kubeconfig)
	if err != nil {
		return nil, err
	}

	// Encapsulate restclient.AddUserAgent to shorten the following code lines
	addAgent := func(config *restclient.Config) *restclient.Config {
		return kubeutil.AddUserAgent(config, ControllerUserAgent)
	}
	c := &controllerconfig.Config{
		Kubeconfig: kubeutil.SetClientOptionsForController(addAgent(kubeconfig)),
	}

	if err := o.ApplyTo(c); err != nil {
		return nil, err
	}

	return c, nil
}
