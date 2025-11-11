package cmd

import (
	"github.com/iton0/hkup-cli/internal/logic/config"
	"github.com/spf13/cobra"
)

var setConfigCmd = &cobra.Command{
	Use:   "set <config-setting> <value>",
	Short: "Set a HkUp config setting",
	Args:  cobra.ExactArgs(2),
	RunE:  config.Set,
}

func init() {
	configCmd.AddCommand(setConfigCmd)
}
