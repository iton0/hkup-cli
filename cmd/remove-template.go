package cmd

import (
	"github.com/iton0/hkup-cli/internal/logic"
	"github.com/spf13/cobra"
)

var removeTemplateCmd = &cobra.Command{
	Use:   "remove <template-name>",
	Short: "Remove a git hook template",
	Args:  cobra.ExactArgs(1),
	RunE:  logic.RemoveTemplate,
}

func init() {
	templateCmd.AddCommand(removeTemplateCmd)
}
