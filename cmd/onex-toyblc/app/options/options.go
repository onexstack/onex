// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

// Package options contains flags and options for initializing an apiserver
package options

import (
	"fmt"
	"time"

	"github.com/onexstack/onexstack/pkg/app"
	"github.com/onexstack/onexstack/pkg/log"
	genericoptions "github.com/onexstack/onexstack/pkg/options"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	cliflag "k8s.io/component-base/cli/flag"

	"github.com/onexstack/onex/internal/pkg/zflag"
	"github.com/onexstack/onex/internal/toyblc"
	"github.com/onexstack/onex/internal/toyblc/pkg/defaults"
)

const (
	// UserAgent is the userAgent name when starting onex-gateway server.
	UserAgent = "onex-toyblc"
)

// ServerOptions contains the configuration options for the server.
type ServerOptions struct {
	Miner            bool                        `json:"miner" mapstructure:"miner"`
	MinMineInterval  time.Duration               `json:"min-mine-interval" mapstructure:"min-mine-interval"`
	MiningDifficulty int                         `json:"mining-difficulty" mapstructure:"mining-difficulty"`
	Address          string                      `json:"address" mapstructure:"address"`
	Accounts         map[string]string           `json:"accounts" mapstructure:"-"`
	P2PAddr          string                      `json:"p2p-addr" mapstructure:"p2p-addr"`
	Peers            []string                    `json:"peers" mapstructure:"peers"`
	HTTPOptions      *genericoptions.HTTPOptions `json:"http" mapstructure:"http"` // blc server
	TLSOptions       *genericoptions.TLSOptions  `json:"tls" mapstructure:"tls"`
	Log              *log.Options                `json:"log" mapstructure:"log"`
}

// Ensure ServerOptions implements the app.CliOptions interface.
var _ app.CliOptions = (*ServerOptions)(nil)

// NewServerOptions creates a ServerOptions instance with default values.
func NewServerOptions() *ServerOptions {
	o := &ServerOptions{
		MinMineInterval:  2 * time.Hour,
		MiningDifficulty: 1,
		Address:          defaults.GenesisAddress,
		Accounts:         defaults.Accounts,
		P2PAddr:          "0.0.0.0:6001",
		Peers:            []string{"ws://localhost:6001"},
		HTTPOptions:      genericoptions.NewHTTPOptions(),
		TLSOptions:       genericoptions.NewTLSOptions(),
		Log:              log.NewOptions(),
	}

	return o
}

// Flags returns flags for a specific server by section name.
func (o *ServerOptions) Flags() (fss cliflag.NamedFlagSets) {
	o.HTTPOptions.AddFlags(fss.FlagSet("http"))
	o.TLSOptions.AddFlags(fss.FlagSet("tls"))
	o.Log.AddFlags(fss.FlagSet("log"))

	// Note: the weird ""+ in below lines seems to be the only way to get gofmt to
	// arrange these text blocks sensibly. Grrr.
	fs := fss.FlagSet("misc")
	fs.BoolVar(&o.Miner, "miner", o.Miner, "Turn on mining mode.")
	fs.DurationVar(&o.MinMineInterval, "min-mine-interval", o.MinMineInterval, "Specify the minimum mining interval.")
	fs.IntVar(&o.MiningDifficulty, "mining-difficulty", o.MiningDifficulty, "Specify the mining difficulty.")
	fs.StringVar(&o.Address, "address", o.Address, "Wallet account to receive the block rewards.")
	fs.StringVar(&o.P2PAddr, "p2p-addr", o.P2PAddr, "The p2p server address.")
	zflag.MapVar(&o.Accounts, "accounts", o.Accounts, "Authentication username and password set for API interface.", fs)
	fs.StringSliceVar(&o.Peers, "peers", o.Peers, "The initial peers.")

	return fss
}

// Complete completes all the required options.
func (o *ServerOptions) Complete() error {
	return nil
}

// Validate checks whether the options in ServerOptions are valid.
func (o *ServerOptions) Validate() error {
	errs := []error{}

	if o.MiningDifficulty < 0 {
		errs = append(errs, fmt.Errorf("`--mining-difficulty` must be non-negative"))
	}

	if err := genericoptions.ValidateAddress(o.P2PAddr); err != nil {
		errs = append(errs, err)
	}

	if len(o.Accounts) == 0 {
		errs = append(errs, fmt.Errorf("empty list of authorized credentials"))
	} else {
		for user := range o.Accounts {
			if user == "" {
				errs = append(errs, fmt.Errorf("account username can not be empty"))
			}
		}
	}

	errs = append(errs, o.HTTPOptions.Validate()...)
	errs = append(errs, o.TLSOptions.Validate()...)
	errs = append(errs, o.Log.Validate()...)

	return utilerrors.NewAggregate(errs)
}

// Config builds an toyblc.Config based on ServerOptions.
func (o *ServerOptions) Config() (*toyblc.Config, error) {
	return &toyblc.Config{
		Miner:           o.Miner,
		MinMineInterval: o.MinMineInterval,
		Address:         o.Address,
		Accounts:        o.Accounts,
		HTTPOptions:     o.HTTPOptions,
		TLSOptions:      o.TLSOptions,
		P2PAddr:         o.P2PAddr,
		Peers:           o.Peers,
	}, nil
}
