package cmd

import (
	"github.com/iton0/hkup-cli/internal/git"
	"github.com/iton0/hkup-cli/internal/logic"
	"github.com/iton0/hkup-cli/internal/util"
	"github.com/spf13/cobra"
)

var (
	addCmd = &cobra.Command{
		Use:       "add <hook-name>",
		Short:     "Add git hook",
		ValidArgs: util.ConvertMapKeysToSlice(git.Hooks()),
		Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		RunE:      logic.Add,
	}
)

func init() {
	addCmd.Flags().StringVar(&logic.Lang, "lang", "", "supported languages for git hooks")
	rootCmd.AddCommand(addCmd)
}
