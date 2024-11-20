package cmd

import (
	"github.com/iton0/hkup-cli/internal/git"
	"github.com/iton0/hkup-cli/internal/logic"
	"github.com/iton0/hkup-cli/internal/util"
	"github.com/spf13/cobra"
)

var (
	removeCmd = &cobra.Command{
		Use:       "remove <hook-name>",
		Aliases:   []string{"rm"},
		Short:     "Remove git hook",
		ValidArgs: util.ConvertMapKeysToSlice(git.Hooks()),
		Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		RunE:      logic.Remove,
	}
)

func init() {}
