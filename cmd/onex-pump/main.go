// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package main

import (
	// Import the automaxprocs package, which automatically configures the GOMAXPROCS
	// value at program startup to match the Linux container's CPU quota.
	// This avoids performance issues caused by an inappropriate default GOMAXPROCS
	// value when running in containers, ensuring that the Go program can fully utilize
	// available CPU resources and avoid CPU waste.
	_ "go.uber.org/automaxprocs/maxprocs"

	"github.com/onexstack/onex/cmd/onex-pump/app"
)

func main() {
	// Creating a new instance of the pump application and running it.
	app.NewApp().Run()
}
