package main

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/giantswarm/microerror"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
)

func main() {
	err := mainE()
	if err != nil {
		panic(err)
	}
}

func mainE() error {
	log.SetOutput(io.Discard)

	start := graphql.Now()

	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		return microerror.Maskf(cannotLoadConfigError, "%s", err.Error())
	}

	err = api.Generate(cfg)
	if err != nil {
		return microerror.Maskf(cannotGenerateCodeError, "%s", err.Error())
	}

	fmt.Printf("Generated %s in %4.2fs\n", cfg.Exec.ImportPath(), time.Since(start).Seconds())

	return nil
}
