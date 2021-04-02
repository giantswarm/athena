package main

import (
	"context"

	"github.com/giantswarm/microerror"
	"github.com/spf13/cobra"

	"github.com/giantswarm/athena/cmd"
)

func main() {
	err := mainE(context.Background())
	if err != nil {
		panic(err)
	}
}

func mainE(ctx context.Context) error {
	var err error

	var rootCommand *cobra.Command
	{
		c := cmd.Config{}

		rootCommand, err = cmd.New(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	err = rootCommand.ExecuteContext(ctx)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
