// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package app

import (
	"github.com/spf13/pflag"
	cliflag "k8s.io/component-base/cli/flag"
)

// OptionsValidator provides methods to complete and validate options.
// Any component requiring options validation should implement this interface.
type OptionsValidator interface {
	// Complete completes all the required options.
	Complete() error

	// Validate validates all the required options.
	Validate() error
}

// NamedFlagSetOptions provides access to server-specific flag sets and embeds the
// validation functionality.
type NamedFlagSetOptions interface {
	// Flags returns flags for a specific server by section name.
	Flags() cliflag.NamedFlagSets

	OptionsValidator
}

// FlagSetOptions defines an interface for command-line options that can
// add themselves to a flag set and perform validation.
type FlagSetOptions interface {
	// AddFlags adds command-specific flags to the provided flag set.
	AddFlags(fs *pflag.FlagSet)

	OptionsValidator
}
