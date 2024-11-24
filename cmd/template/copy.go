package template

import (
	"github.com/iton0/hkup-cli/internal/logic/template"
	"github.com/spf13/cobra"
)

var copyCmd = &cobra.Command{
	Use:   "copy <template-name>",
	Short: "Copy a git hook template",
	Args:  cobra.ExactArgs(1),
	RunE:  template.Copy,
}

func init() {}
