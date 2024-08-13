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

	"github.com/giantswarm/athena/internal/config"
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
		setViperDefaults(r.viper, cmd)

		err = r.readConfig(cmd, r.flag.ConfigDir, r.flag.ConfigFile)
		if err != nil {
			return microerror.Mask(err)
		}

		err = r.readConfig(cmd, r.flag.SecretDir, r.flag.SecretFile)
		if err != nil {
			return microerror.Mask(err)
		}

		f := config.New()
		err = r.viper.Unmarshal(f)
		if err != nil {
			return microerror.Mask(err)
		}

		// Set default values.
		r.viper.SetDefault("server.listenAddress", ":8000")
		r.viper.SetDefault("server.allowedOrigins", []string{"*"})
	}

	var s *server.Server
	{
		c := server.Config{
			Log:                      r.log,
			AllowedOrigins:           r.viper.GetStringSlice("server.allowedOrigins"),
			ListenAddress:            r.viper.GetString("server.listenAddress"),
			EnableIntrospection:      r.viper.GetBool("server.enableIntrospection"),
			InstallationProvider:     r.viper.GetString("identity.provider"),
			InstallationCodename:     r.viper.GetString("identity.codename"),
			InstallationK8sApiUrl:    r.viper.GetString("kubernetes.apiUrl"),
			InstallationK8sAuthUrl:   r.viper.GetString("kubernetes.authUrl"),
			InstallationK8sCaCert:    r.viper.GetString("kubernetes.caCert"),
			AnalyticsEnv:             r.viper.GetString("analytics.environmentType"),
			AnalyticsCredentialsJSON: r.viper.GetString("analytics.credentialsJSON"),
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

func (r *runner) readConfig(cmd *cobra.Command, dir string, file string) error {
	v := viper.New()

	setViperDefaults(v, cmd)

	v.AddConfigPath(dir)
	v.SetConfigName(file)

	err := v.ReadInConfig()
	if errors.As(err, &viper.ConfigFileNotFoundError{}) {
		// Ignore error, we can use defaults.
	} else if err != nil {
		return microerror.Mask(err)
	}

	for _, k := range v.AllKeys() {
		r.viper.Set(k, v.Get(k))
	}

	return nil
}

func setViperDefaults(viperInstance *viper.Viper, cmd *cobra.Command) {
	viperInstance.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viperInstance.SetEnvPrefix(cmd.Root().Name())
	viperInstance.AutomaticEnv()
}
