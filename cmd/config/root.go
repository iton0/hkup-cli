package config

import "github.com/spf13/cobra"

// RootCmd is the root config subcommand that will be added to the root HkUp
// command.
var RootCmd = &cobra.Command{
	Use:    "config",
	Short:  "HkUp configuration settings",
	Hidden: true, // TODO: remove after finalizing configuration settings
}

func init() {
	RootCmd.AddCommand(getCmd)
	RootCmd.AddCommand(setCmd)
}
