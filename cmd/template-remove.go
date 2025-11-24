package cmd

import (
	"github.com/iton0/hkup-cli/v2/internal/logic"
	"github.com/spf13/cobra"
)

var templateRemoveCmd = &cobra.Command{
	Use:   "remove <template-name>",
	Short: "Remove a git hook template",
	Args:  cobra.ExactArgs(1),
	RunE:  logic.TemplateRemove,
}

func init() {
	templateCmd.AddCommand(templateRemoveCmd)
}
