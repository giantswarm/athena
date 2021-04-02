package daemon

import (
	"context"
	"errors"
	"io"
	"strings"

	"github.com/giantswarm/microerror"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	envFlags "github.com/giantswarm/athena/internal/flags"
)

type runner struct {
	flag  *flag
	log   *zap.SugaredLogger
	viper *viper.Viper

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
	var err error

	{
		r.viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		r.viper.SetEnvPrefix(cmd.Root().Name())

		r.viper.AddConfigPath(r.flag.ConfigDir)
		r.viper.SetConfigName(r.flag.ConfigFile)

		r.viper.AutomaticEnv()

		err = r.viper.ReadInConfig()
		if errors.As(err, &viper.ConfigFileNotFoundError{}) {
			// Ignore error, we can use defaults.
		} else if err != nil {
			return microerror.Mask(err)
		}

		f := envFlags.New()
		err = r.viper.Unmarshal(f)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}
