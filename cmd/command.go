package cmd

import (
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/giantswarm/microerror"

	"github.com/giantswarm/athena/cmd/daemon"
	"github.com/giantswarm/athena/pkg/project"
)

type Config struct {
	Stderr io.Writer
	Stdout io.Writer
}

func New(config Config) (*cobra.Command, error) {
	if config.Stderr == nil {
		config.Stderr = os.Stderr
	}
	if config.Stdout == nil {
		config.Stdout = os.Stdout
	}

	var err error

	var daemonCmd *cobra.Command
	{
		c := daemon.Config{
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		daemonCmd, err = daemon.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	f := &flag{}

	r := &runner{
		flag:   f,
		stderr: config.Stderr,
		stdout: config.Stdout,
	}

	c := &cobra.Command{
		Use:          project.Name(),
		Short:        project.Description(),
		Long:         project.Description(),
		RunE:         r.Run,
		SilenceUsage: true,
		Version:      project.Version(),
	}

	f.Init(c)

	c.AddCommand(daemonCmd)

	return c, nil
}
