package cmd

import (
	"github.com/iton0/hkup-cli/v2/internal/logic"
	"github.com/spf13/cobra"
)

var templateCopyCmd = &cobra.Command{
	Use:   "copy <template-name>",
	Short: "Copy a git hook template",
	Args:  cobra.ExactArgs(1),
	RunE:  logic.TemplateCopy,
}

func init() {
	templateCmd.AddCommand(templateCopyCmd)
}
