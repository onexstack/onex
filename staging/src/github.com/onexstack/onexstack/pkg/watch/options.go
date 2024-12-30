package watch

import (
	"errors"

	"github.com/spf13/pflag"
)

// Options structure holds the configuration options required to create and run a watch server.
type Options struct {
	// DisableWatchers is a slice of watchers that will be disabled when the server is run.
	DisableWatchers []string `json:"disable-watchers" mapstructure:"disable-watchers"`

	// MaxWorkers defines the maximum number of concurrent workers that each watcher can spawn.
	MaxWorkers int64 `json:"max-workers" mapstructure:"max-workers"`
}

// NewOptions initializes and returns a new Options instance with default values.
func NewOptions() *Options {
	o := &Options{
		DisableWatchers: []string{},
		MaxWorkers:      10,
	}

	return o
}

// AddFlags adds the command-line flags associated with the Options structure to the provided FlagSet.
// This will allow users to configure the watch server via command-line arguments.
func (o *Options) AddFlags(fs *pflag.FlagSet) {
	fs.StringSliceVar(&o.DisableWatchers, "disable-watchers", o.DisableWatchers, "The list of watchers that should be disabled.")
	fs.Int64Var(&o.MaxWorkers, "max-workers", o.MaxWorkers, "Specify the maximum concurrency worker of each watcher.")
}

// Validate checks the Options structure for required configurations and returns a slice of errors.
func (o *Options) Validate() []error {
	errs := []error{}

	if o.MaxWorkers <= 0 {
		errs = append(errs, errors.New("max-workers must be greater than 0"))
	}

	return errs
}
