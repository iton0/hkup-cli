package cmd

import (
	"github.com/iton0/hkup-cli/internal/logic"
	"github.com/spf13/cobra"
)

var configSetCmd = &cobra.Command{
	Use:   "set <config-setting> <value>",
	Short: "Set a HkUp config setting",
	Args:  cobra.ExactArgs(2),
	RunE:  logic.ConfigSet,
}

func init() {
	configCmd.AddCommand(configSetCmd)
}
