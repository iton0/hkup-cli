package template

import (
	"github.com/iton0/hkup-cli/internal/logic/template"
	"github.com/spf13/cobra"
)

var (
	editCmd = &cobra.Command{
		Use:   "edit <template-name>",
		Short: "Edit a git hook template",
		Args:  cobra.ExactArgs(1),
		RunE:  template.Edit,
	}
)

func init() {}
