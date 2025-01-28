// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

// Package options contains flags and options for initializing an apiserver
package options

import (
	"github.com/onexstack/onexstack/pkg/app"
	"github.com/onexstack/onexstack/pkg/log"
	genericoptions "github.com/onexstack/onexstack/pkg/options"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/client-go/tools/clientcmd"
	cliflag "k8s.io/component-base/cli/flag"

	"github.com/onexstack/onex/internal/gateway"
	"github.com/onexstack/onex/internal/pkg/client"
	"github.com/onexstack/onex/internal/pkg/client/usercenter"
	"github.com/onexstack/onex/internal/pkg/feature"
	kubeutil "github.com/onexstack/onex/internal/pkg/util/kube"
)

const (
	// UserAgent is the userAgent name when starting onex-gateway server.
	UserAgent = "onex-gateway"
)

// ServerOptions contains the configuration options for the server.
type ServerOptions struct {
	// GenericOptions *genericoptions.Options       `json:"server"   mapstructure:"server"`
	GRPCOptions       *genericoptions.GRPCOptions    `json:"grpc" mapstructure:"grpc"`
	HTTPOptions       *genericoptions.HTTPOptions    `json:"http" mapstructure:"http"`
	TLSOptions        *genericoptions.TLSOptions     `json:"tls" mapstructure:"tls"`
	MySQLOptions      *genericoptions.MySQLOptions   `json:"mysql" mapstructure:"mysql"`
	RedisOptions      *genericoptions.RedisOptions   `json:"redis" mapstructure:"redis"`
	EtcdOptions       *genericoptions.EtcdOptions    `json:"etcd" mapstructure:"etcd"`
	JaegerOptions     *genericoptions.JaegerOptions  `json:"jaeger" mapstructure:"jaeger"`
	ConsulOptions     *genericoptions.ConsulOptions  `json:"consul" mapstructure:"consul"`
	UserCenterOptions *usercenter.UserCenterOptions  `json:"usercenter" mapstructure:"usercenter"`
	MetricsOptions    *genericoptions.MetricsOptions `json:"metrics" mapstructure:"metrics"`
	EnableTLS         bool                           `json:"enable-tls" mapstructure:"enable-tls"`
	// Path to kubeconfig file with authorization and master location information.
	Kubeconfig   string          `json:"kubeconfig" mapstructure:"kubeconfig"`
	FeatureGates map[string]bool `json:"feature-gates"`

	Log *log.Options `json:"log" mapstructure:"log"`
}

// Ensure ServerOptions implements the app.CliOptions interface.
var _ app.CliOptions = (*ServerOptions)(nil)

// NewServerOptions creates a ServerOptions instance with default values.
func NewServerOptions() *ServerOptions {
	o := &ServerOptions{
		// GenericOptions: genericoptions.NewOptions(),
		GRPCOptions:       genericoptions.NewGRPCOptions(),
		HTTPOptions:       genericoptions.NewHTTPOptions(),
		TLSOptions:        genericoptions.NewTLSOptions(),
		MySQLOptions:      genericoptions.NewMySQLOptions(),
		RedisOptions:      genericoptions.NewRedisOptions(),
		EtcdOptions:       genericoptions.NewEtcdOptions(),
		JaegerOptions:     genericoptions.NewJaegerOptions(),
		ConsulOptions:     genericoptions.NewConsulOptions(),
		UserCenterOptions: usercenter.NewUserCenterOptions(),
		MetricsOptions:    genericoptions.NewMetricsOptions(),
		Log:               log.NewOptions(),
	}

	return o
}

// Flags returns flags for a specific server by section name.
func (o *ServerOptions) Flags() (fss cliflag.NamedFlagSets) {
	o.GRPCOptions.AddFlags(fss.FlagSet("grpc"))
	o.HTTPOptions.AddFlags(fss.FlagSet("http"))
	o.TLSOptions.AddFlags(fss.FlagSet("tls"))
	o.MySQLOptions.AddFlags(fss.FlagSet("mysql"))
	o.RedisOptions.AddFlags(fss.FlagSet("redis"))
	o.EtcdOptions.AddFlags(fss.FlagSet("etcd"))
	o.JaegerOptions.AddFlags(fss.FlagSet("jaeger"))
	o.ConsulOptions.AddFlags(fss.FlagSet("consul"))
	o.UserCenterOptions.AddFlags(fss.FlagSet("usercenter"))
	o.MetricsOptions.AddFlags(fss.FlagSet("metrics"))
	o.Log.AddFlags(fss.FlagSet("log"))

	// Note: the weird ""+ in below lines seems to be the only way to get gofmt to
	// arrange these text blocks sensibly. Grrr.
	fs := fss.FlagSet("misc")
	client.AddFlags(fs)
	fs.StringVar(&o.Kubeconfig, "kubeconfig", o.Kubeconfig, "Path to kubeconfig file with authorization and master location information.")
	feature.DefaultMutableFeatureGate.AddFlag(fs)

	return fss
}

// Complete completes all the required options.
func (o *ServerOptions) Complete() error {
	if o.JaegerOptions.ServiceName == "" {
		o.JaegerOptions.ServiceName = UserAgent
	}
	_ = feature.DefaultMutableFeatureGate.SetFromMap(o.FeatureGates)
	return nil
}

// Validate checks whether the options in ServerOptions are valid.
func (o *ServerOptions) Validate() error {
	errs := []error{}

	errs = append(errs, o.GRPCOptions.Validate()...)
	errs = append(errs, o.HTTPOptions.Validate()...)
	errs = append(errs, o.TLSOptions.Validate()...)
	errs = append(errs, o.MySQLOptions.Validate()...)
	errs = append(errs, o.RedisOptions.Validate()...)
	errs = append(errs, o.EtcdOptions.Validate()...)
	errs = append(errs, o.JaegerOptions.Validate()...)
	errs = append(errs, o.ConsulOptions.Validate()...)
	errs = append(errs, o.UserCenterOptions.Validate()...)
	errs = append(errs, o.MetricsOptions.Validate()...)
	errs = append(errs, o.Log.Validate()...)

	return utilerrors.NewAggregate(errs)
}

// Config builds an gateway.Config based on ServerOptions.
func (o *ServerOptions) Config() (*gateway.Config, error) {
	kubeconfig, err := clientcmd.BuildConfigFromFlags("", o.Kubeconfig)
	if err != nil {
		return nil, err
	}
	kubeconfig = kubeutil.SetDefaultClientOptions(kubeutil.AddUserAgent(kubeconfig, UserAgent))
	cfg := &gateway.Config{
		GRPCOptions:       o.GRPCOptions,
		HTTPOptions:       o.HTTPOptions,
		TLSOptions:        o.TLSOptions,
		UserCenterOptions: o.UserCenterOptions,
		MySQLOptions:      o.MySQLOptions,
		RedisOptions:      o.RedisOptions,
		EtcdOptions:       o.EtcdOptions,
		JaegerOptions:     o.JaegerOptions,
		ConsulOptions:     o.ConsulOptions,
		Kubeconfig:        kubeconfig,
	}

	return cfg, nil
}
