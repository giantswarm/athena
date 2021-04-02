package daemon

import (
	"context"
	"io"

	"github.com/giantswarm/microerror"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	globalFlags "github.com/giantswarm/athena/internal/flags"
)

type runner struct {
	flag   *flag
	log    *zap.SugaredLogger
	stdout io.Writer
	stderr io.Writer
}

func (r *runner) Run(cmd *cobra.Command, args []string) error {
	err := r.flag.Validate()
	if err != nil {
		return microerror.Mask(err)
	}

	err = r.run(cmd.Context(), cmd, args)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

func (r *runner) run(ctx context.Context, cmd *cobra.Command, args []string) error {
	f := globalFlags.New()
	{
		f.Address = r.flag.Address
		f.AllowedOrigins = r.flag.AllowedOrigins
	}

	return nil
}
