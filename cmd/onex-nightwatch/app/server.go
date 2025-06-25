// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/onexstack/onexstack/pkg/app"
	genericoptions "github.com/onexstack/onexstack/pkg/options"
	"gorm.io/gorm"
	genericapiserver "k8s.io/apiserver/pkg/server"

	"github.com/onexstack/onex/cmd/onex-nightwatch/app/options"
	"github.com/onexstack/onex/internal/nightwatch"
)

const commandDesc = `The nightwatch server is responsible for executing some async tasks 
like linux cronjob. You can add Cron(github.com/robfig/cron) jobs on the given schedule
use the Cron spec format.`

// jobServer represents the HTTP server with optional TLS and graceful shutdown capabilities.
type jobServer struct {
	stopCh     <-chan struct{}
	tlsOptions *genericoptions.TLSOptions
}

// Option is a function that configures the jobServer.
type Option func(jrs *jobServer)

// WithTLSOptions sets the TLS options for the job REST server.
func WithTLSOptions(tlsOptions *genericoptions.TLSOptions) Option {
	return func(jrs *jobServer) {
		jrs.tlsOptions = tlsOptions
	}
}

// WithStopChannel sets the stop channel for graceful shutdown.
func WithStopChannel(stopCh <-chan struct{}) Option {
	return func(jrs *jobServer) {
		jrs.stopCh = stopCh
	}
}

// NewApp creates and returns a new App object with default parameters.
func NewApp(appName string) *app.App {
	opts := options.NewServerOptions()
	application := app.NewApp(
		appName,
		"Launch an asynchronous task processing server",
		app.WithDescription(commandDesc),
		app.WithOptions(opts),
		app.WithDefaultValidArgs(),
		app.WithRunFunc(run(opts)),
	)

	return application
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

// NewJobServer creates a new instance of the job server with the specified options.
func NewJobServer(httpOptions *genericoptions.HTTPOptions, db *gorm.DB, opts ...Option) *nightwatch.RESTServer {
	jrs := jobServer{}
	for _, opt := range opts {
		opt(&jrs)
	}

	return nightwatch.NewRESTServer(httpOptions, jrs.tlsOptions, db)
}

// InstallJobAPI sets up the job-related routes in the provided router.
func InstallJobAPI(engine *gin.Engine, db *gorm.DB) {
	nightwatch.InstallJobAPI(engine, db)
}
