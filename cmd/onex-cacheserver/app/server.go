package app

import (
	"fmt"

	"github.com/onexstack/onexstack/pkg/app"
	genericapiserver "k8s.io/apiserver/pkg/server"

	"github.com/onexstack/onex/cmd/onex-cacheserver/app/options"
)

const commandDesc = `onex-cacheserver is an example cache server, demonstrating   
how to develop a caching service.  

You can find more information about onex-cacheserver at:  
https://github.com/onexstack/onex/blob/master/docs/guide/en-US/cmd/onex-cacheserver.md  
`

// NewApp creates and returns a new App object with default parameters.
func NewApp() *app.App {
	opts := options.NewServerOptions()

	return app.NewApp(
		"onex-cacheserver",               // Application name
		"Launch an onex cache server",    // Brief description
		app.WithDescription(commandDesc), // Full descriptive documentation
		app.WithOptions(opts),            // Application-specific options
		app.WithDefaultValidArgs(),       // Command-line argument validation setup
		app.WithRunFunc(run(opts)),       // The run function for starting the app
	)
}

// run contains the main logic for initializing and running the server.
func run(opts *options.ServerOptions) app.RunFunc {
	return func() error {
		// Load the configuration options
		cfg, err := opts.Config()
		if err != nil {
			return fmt.Errorf("failed to load configuration: %w", err)
		}

		ctx := genericapiserver.SetupSignalContext()

		// Build the server using the configuration
		server, err := cfg.NewServer(ctx)
		if err != nil {
			return fmt.Errorf("failed to create server: %w", err)
		}

		// Run the server with signal context for graceful shutdown
		return server.Run(ctx)
	}
}
