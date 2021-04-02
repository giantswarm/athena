package cmd

import (
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/giantswarm/microerror"

	"github.com/giantswarm/athena/cmd/daemon"
	"github.com/giantswarm/athena/pkg/project"
)

type Config struct {
	Log   *zap.SugaredLogger
	Viper *viper.Viper

	Stderr io.Writer
	Stdout io.Writer
}

func New(config Config) (*cobra.Command, error) {
	if config.Log == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Log must not be empty", config)
	}
	if config.Viper == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Viper must not be empty", config)
	}
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
			Log:    config.Log,
			Viper:  config.Viper,
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
		log:    config.Log,
		viper:  config.Viper,
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
