package daemon

import (
	"github.com/spf13/cobra"
)

const (
	configDir  = "config-dir"
	configFile = "config-file"
	secretDir  = "secret-dir"
	secretFile = "secret-file"
)

type flag struct {
	ConfigDir  string
	ConfigFile string
	SecretDir  string
	SecretFile string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&f.ConfigDir, configDir, ".", "")
	cmd.PersistentFlags().StringVar(&f.ConfigFile, configFile, "config", "")
	cmd.PersistentFlags().StringVar(&f.SecretDir, secretDir, ".", "")
	cmd.PersistentFlags().StringVar(&f.SecretFile, secretFile, "secret", "")
}

func (f *flag) Validate() error {
	return nil
}
