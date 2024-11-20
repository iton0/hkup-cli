package cmd

import (
	"github.com/iton0/hkup-cli/internal/logic"
	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:       "list {hook|lang}",
		Aliases:   []string{"ls"},
		Short:     "List git hooks information",
		ValidArgs: []string{"hook", "lang"},
		Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		RunE:      logic.List,
	}
)

func init() {}
