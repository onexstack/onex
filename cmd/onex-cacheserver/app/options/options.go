// Package options provides flags and configuration for initializing the Onex Cache Server.
package options

import (
	"github.com/onexstack/onexstack/pkg/app"
	"github.com/onexstack/onexstack/pkg/log"
	genericoptions "github.com/onexstack/onexstack/pkg/options"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	cliflag "k8s.io/component-base/cli/flag"

	"github.com/onexstack/onex/internal/cacheserver"
)

// UserAgent is the default name used for the Onex Cache Server client.
const UserAgent = "onex-cacheserver"

// ServerOptions contains the configuration options for the server.
type ServerOptions struct {
	DisableCache   bool                           `json:"disable-cache" mapstructure:"disable-cache"`
	GRPCOptions    *genericoptions.GRPCOptions    `json:"grpc" mapstructure:"grpc"`
	TLSOptions     *genericoptions.TLSOptions     `json:"tls" mapstructure:"tls"`
	RedisOptions   *genericoptions.RedisOptions   `json:"redis" mapstructure:"redis"`
	MySQLOptions   *genericoptions.MySQLOptions   `json:"mysql" mapstructure:"mysql"`
	JaegerOptions  *genericoptions.JaegerOptions  `json:"jaeger" mapstructure:"jaeger"`
	MetricsOptions *genericoptions.MetricsOptions `json:"metrics" mapstructure:"metrics"`
	Logging        *log.Options                   `json:"log" mapstructure:"log"`
}

// Ensure ServerOptions implements the app.CliOptions interface.
var _ app.CliOptions = (*ServerOptions)(nil)

// NewServerOptions creates a ServerOptions instance with default values.
func NewServerOptions() *ServerOptions {
	o := &ServerOptions{
		DisableCache:   false,
		GRPCOptions:    genericoptions.NewGRPCOptions(),
		TLSOptions:     genericoptions.NewTLSOptions(),
		RedisOptions:   genericoptions.NewRedisOptions(),
		MySQLOptions:   genericoptions.NewMySQLOptions(),
		JaegerOptions:  genericoptions.NewJaegerOptions(),
		MetricsOptions: genericoptions.NewMetricsOptions(),
		Logging:        log.NewOptions(),
	}

	return o
}

// Flags returns flags for a specific server by section name.
func (o *ServerOptions) Flags() (fss cliflag.NamedFlagSets) {
	// Add flags for each option group with meaningful section names.
	o.GRPCOptions.AddFlags(fss.FlagSet("grpc"))
	o.TLSOptions.AddFlags(fss.FlagSet("tls"))
	o.RedisOptions.AddFlags(fss.FlagSet("redis"))
	o.MySQLOptions.AddFlags(fss.FlagSet("mysql"))
	o.JaegerOptions.AddFlags(fss.FlagSet("jaeger"))
	o.MetricsOptions.AddFlags(fss.FlagSet("metrics"))
	o.Logging.AddFlags(fss.FlagSet("log"))

	// Add a miscellaneous flag for the cache control feature.
	miscFs := fss.FlagSet("misc")
	miscFs.BoolVar(&o.DisableCache, "disable-cache", o.DisableCache, "Disable the local memory cache.")

	return fss
}

// Complete completes all the required options.
func (o *ServerOptions) Complete() error {
	if o.JaegerOptions.ServiceName == "" {
		o.JaegerOptions.ServiceName = UserAgent
	}
	return nil
}

// Validate checks whether the options in ServerOptions are valid.
func (o *ServerOptions) Validate() error {
	var errs []error

	// Perform validation for each option group, accumulating errors.
	errs = append(errs, o.GRPCOptions.Validate()...)
	errs = append(errs, o.TLSOptions.Validate()...)
	errs = append(errs, o.RedisOptions.Validate()...)
	errs = append(errs, o.MySQLOptions.Validate()...)
	errs = append(errs, o.JaegerOptions.Validate()...)
	errs = append(errs, o.MetricsOptions.Validate()...)
	errs = append(errs, o.Logging.Validate()...)

	// Aggregate all validation errors into a single error object.
	return utilerrors.NewAggregate(errs)
}

// Config builds an cacheserver.Config based on ServerOptions.
func (o *ServerOptions) Config() (*cacheserver.Config, error) {
	// Ensure the configuration includes all relevant fields from the options.
	return &cacheserver.Config{
		DisableCache:  o.DisableCache,
		GRPCOptions:   o.GRPCOptions,
		TLSOptions:    o.TLSOptions,
		RedisOptions:  o.RedisOptions,
		MySQLOptions:  o.MySQLOptions,
		JaegerOptions: o.JaegerOptions,
	}, nil
}
