package cmd

import (
	"github.com/iton0/hkup-cli/v2/internal/logic"
	"github.com/spf13/cobra"
)

var configGetCmd = &cobra.Command{
	Use:   "get <config-setting>",
	Short: "Get a HkUp config setting",
	Args:  cobra.ExactArgs(1),
	RunE:  logic.ConfigGet,
}

func init() {
	configCmd.AddCommand(configGetCmd)
}
