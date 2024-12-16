package cmd

import (
	"github.com/iton0/hkup-cli/internal/logic"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [hook|lang|template]",
	Short: "List git hooks information",
	Long: `List git hooks information for the specified category.

If no arguments are provided, this command will display the hooks used
in the current working directory.

Valid arguments:
- hook:     Displays supported git hooks.
- lang:     Displays supported languages used for hooks.
- template: Displays user-defined templates.`,
	ValidArgs: []string{"hook", "lang", "template"},
	Args:      cobra.MatchAll(cobra.MaximumNArgs(1), cobra.OnlyValidArgs),
	RunE:      logic.List,
}

func init() {}
