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
	cliflag "k8s.io/component-base/cli/flag"

	"github.com/onexstack/onex/internal/pkg/client"
	"github.com/onexstack/onex/internal/pkg/feature"
	known "github.com/onexstack/onex/internal/pkg/known/usercenter"
	"github.com/onexstack/onex/internal/usercenter"
)

const (
	// UserAgent is the userAgent name when starting onex-usercenter server.
	UserAgent = "onex-usercenter"
)

// ServerOptions contains the configuration options for the server.
type ServerOptions struct {
	// GenericOptions *genericoptions.Options       `json:"server"   mapstructure:"server"`
	// gRPC options for configuring gRPC related options.
	GRPCOptions *genericoptions.GRPCOptions `json:"grpc" mapstructure:"grpc"`
	// HTTP options for configuring HTTP related options.
	HTTPOptions *genericoptions.HTTPOptions `json:"http" mapstructure:"http"`
	// TLS options for configuring TLS related options.
	TLSOptions *genericoptions.TLSOptions `json:"tls" mapstructure:"tls"`
	// MySQL options for configuring MySQL database related options.
	MySQLOptions *genericoptions.MySQLOptions `json:"mysql" mapstructure:"mysql"`
	// Redis options for configuring Redis related options.
	RedisOptions *genericoptions.RedisOptions `json:"redis" mapstructure:"redis"`
	// Etcd options for configuring Etcd related options.
	EtcdOptions *genericoptions.EtcdOptions `json:"etcd" mapstructure:"etcd"`
	// Kafka options for configuring Kafka related options.
	KafkaOptions *genericoptions.KafkaOptions `json:"kafka" mapstructure:"kafka"`
	// Jaeger options for configuring Jaeger related options.
	JaegerOptions *genericoptions.JaegerOptions `json:"jaeger" mapstructure:"jaeger"`
	// Consul options for configuring Consul related options.
	ConsulOptions *genericoptions.ConsulOptions `json:"consul" mapstructure:"consul"`
	// JWT options for configuring JWT related options.
	JWTOptions *genericoptions.JWTOptions `json:"jwt" mapstructure:"jwt"`
	// Metrics options for configuring metric related options.
	Metrics *genericoptions.MetricsOptions `json:"metrics" mapstructure:"metrics"`
	// TODO: add `mapstructure` tag for FeatureGates
	// A map of string to boolean values representing feature gates for enabling or disabling specific features.
	FeatureGates map[string]bool `json:"feature-gates"`
	// Log options for configuring log related options.
	Log *log.Options `json:"log" mapstructure:"log"`
}

// Ensure ServerOptions implements the app.CliOptions interface.
var _ app.CliOptions = (*ServerOptions)(nil)

// NewServerOptions creates a ServerOptions instance with default values.
func NewServerOptions() *ServerOptions {
	o := &ServerOptions{
		// GenericOptions: genericoptions.NewOptions(),
		GRPCOptions:   genericoptions.NewGRPCOptions(),
		HTTPOptions:   genericoptions.NewHTTPOptions(),
		TLSOptions:    genericoptions.NewTLSOptions(),
		MySQLOptions:  genericoptions.NewMySQLOptions(),
		RedisOptions:  genericoptions.NewRedisOptions(),
		EtcdOptions:   genericoptions.NewEtcdOptions(),
		KafkaOptions:  genericoptions.NewKafkaOptions(),
		JaegerOptions: genericoptions.NewJaegerOptions(),
		ConsulOptions: genericoptions.NewConsulOptions(),
		JWTOptions:    genericoptions.NewJWTOptions(),
		Metrics:       genericoptions.NewMetricsOptions(),
		Log:           log.NewOptions(),
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
	o.KafkaOptions.AddFlags(fss.FlagSet("kafka"))
	o.JaegerOptions.AddFlags(fss.FlagSet("jaeger"))
	o.ConsulOptions.AddFlags(fss.FlagSet("consul"))
	o.JWTOptions.AddFlags(fss.FlagSet("jwt"))
	o.Metrics.AddFlags(fss.FlagSet("metrics"))
	o.Log.AddFlags(fss.FlagSet("log"))

	// Note: the weird ""+ in below lines seems to be the only way to get gofmt to
	// arrange these text blocks sensibly. Grrr.
	fs := fss.FlagSet("misc")
	client.AddFlags(fs)
	feature.DefaultMutableFeatureGate.AddFlag(fs)

	return fss
}

// Complete completes all the required options.
func (o *ServerOptions) Complete() error {
	if o.JaegerOptions.ServiceName == "" {
		o.JaegerOptions.ServiceName = UserAgent
	}
	o.JWTOptions.Expired = known.RefreshTokenExpire
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
	errs = append(errs, o.KafkaOptions.Validate()...)
	errs = append(errs, o.JaegerOptions.Validate()...)
	errs = append(errs, o.ConsulOptions.Validate()...)
	errs = append(errs, o.JWTOptions.Validate()...)
	errs = append(errs, o.Metrics.Validate()...)
	errs = append(errs, o.Log.Validate()...)

	return utilerrors.NewAggregate(errs)
}

// Config builds an usercenter.Config based on ServerOptions.
func (o *ServerOptions) Config() (*usercenter.Config, error) {
	return &usercenter.Config{
		GRPCOptions:   o.GRPCOptions,
		HTTPOptions:   o.HTTPOptions,
		TLSOptions:    o.TLSOptions,
		JWTOptions:    o.JWTOptions,
		MySQLOptions:  o.MySQLOptions,
		RedisOptions:  o.RedisOptions,
		EtcdOptions:   o.EtcdOptions,
		KafkaOptions:  o.KafkaOptions,
		JaegerOptions: o.JaegerOptions,
		ConsulOptions: o.ConsulOptions,
	}, nil
}
