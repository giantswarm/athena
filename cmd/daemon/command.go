package daemon

import (
	"io"
	"os"

	"github.com/giantswarm/microerror"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

const (
	name        = "daemon"
	description = "Run the API server daemon."
)

type Config struct {
	Log *zap.SugaredLogger

	Stderr io.Writer
	Stdout io.Writer
}

func New(config Config) (*cobra.Command, error) {
	if config.Log == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	if config.Stderr == nil {
		config.Stderr = os.Stderr
	}
	if config.Stdout == nil {
		config.Stdout = os.Stdout
	}

	f := &flag{}

	r := &runner{
		flag:   f,
		log:    config.Log,
		stderr: config.Stderr,
		stdout: config.Stdout,
	}

	c := &cobra.Command{
		Use:   name,
		Short: description,
		Long:  description,
		RunE:  r.Run,
	}

	f.Init(c)

	return c, nil
}
