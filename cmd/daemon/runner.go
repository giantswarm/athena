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
	"github.com/giantswarm/athena/pkg/server"
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
		// Get configuration.
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

		// Set default values.
		r.viper.SetDefault("server.listenAddress", ":8080")
		r.viper.SetDefault("server.allowedOrigins", "*")
	}

	var s *server.Server
	{
		c := server.Config{
			Log:           r.log,
			ListenAddress: r.viper.GetString("server.listenAddress"),
		}

		s, err = server.New(c)
		if err != nil {
			return microerror.Mask(err)
		}

		err = s.Boot()
		if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}
