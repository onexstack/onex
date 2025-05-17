package main

import (
	"log"

	genericoptions "github.com/onexstack/onexstack/pkg/options"
	"github.com/spf13/pflag"

	"github.com/onexstack/onex/cmd/onex-nightwatch/app"
)

func main() {
	httpOptions := genericoptions.NewHTTPOptions()
	httpOptions.AddFlags(pflag.CommandLine)

	mysqlOptions := genericoptions.NewMySQLOptions()
	mysqlOptions.AddFlags(pflag.CommandLine)

	pflag.Parse()

	db, err := mysqlOptions.NewDB()
	if err != nil {
		log.Fatalf("Failed to new db: %v", err)
	}

	app.NewJobServer(httpOptions, db).RunOrDie()
}
