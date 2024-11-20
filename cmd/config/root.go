package config

import (
	"github.com/spf13/cobra"
)

var (
	// RootCmd is the config command that will be added to the root HkUp command.
	RootCmd = &cobra.Command{
		Use:    "config",
		Short:  "HkUp configuration settings",
		Hidden: true, // TODO: remove after finalizing configuration settings
	}
)

func init() {
	RootCmd.AddCommand(getCmd)
	RootCmd.AddCommand(setCmd)
}
