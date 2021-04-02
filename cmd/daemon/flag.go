package daemon

import (
	"github.com/spf13/cobra"
)

const (
	flagAddress       = "address"
	flagAllowedOrigin = "allowed-origin"
)

type flag struct {
	Address        string
	AllowedOrigins []string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&f.Address, flagAddress, "0.0.0.0:8000", "Set the address that the application will run on.")
	cmd.PersistentFlags().StringSliceVar(&f.AllowedOrigins, flagAllowedOrigin, []string{"*"}, "Set the allowed origin for connections.")
}

func (f *flag) Validate() error {
	return nil
}
