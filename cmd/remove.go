package cmd

import (
	"github.com/iton0/hkup-cli/v2/internal/git"
	"github.com/iton0/hkup-cli/v2/internal/logic"
	"github.com/iton0/hkup-cli/v2/internal/util"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:       "remove <hook-name>",
	Short:     "Remove git hook",
	ValidArgs: util.ConvertMapKeysToSlice(git.Hooks()),
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	RunE:      logic.Remove,
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
