package cmd

import (
	"github.com/iton0/hkup-cli/internal/logic"
	"github.com/spf13/cobra"
)

var editTemplateCmd = &cobra.Command{
	Use:   "edit <template-name>",
	Short: "Edit a git hook template",
	Args:  cobra.ExactArgs(1),
	RunE:  logic.EditTemplate,
}

func init() { templateCmd.AddCommand(editTemplateCmd) }
