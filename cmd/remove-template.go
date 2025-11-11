package cmd

import (
	"github.com/iton0/hkup-cli/internal/logic/template"
	"github.com/spf13/cobra"
)

var removeTemplateCmd = &cobra.Command{
	Use:   "remove <template-name>",
	Short: "Remove a git hook template",
	Args:  cobra.ExactArgs(1),
	RunE:  template.Remove,
}

func init() {
	templateCmd.AddCommand(removeTemplateCmd)
}
