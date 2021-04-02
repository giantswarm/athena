package daemon

import (
	"github.com/spf13/cobra"
)

const (
	configDir  = "config-dir"
	configFile = "config-file"
)

type flag struct {
	ConfigDir  string
	ConfigFile string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&f.ConfigDir, configDir, ".", "")
	cmd.PersistentFlags().StringVar(&f.ConfigFile, configFile, "config", "")
}

func (f *flag) Validate() error {
	return nil
}
