package cmd

import (
	"github.com/iton0/hkup-cli/internal/logic"
	"github.com/spf13/cobra"
)

var (
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize hkup",
		Long:  "Create an empty hkup folder or reinitialize an existing one",
		Args:  cobra.NoArgs,
		RunE:  logic.Init,
	}
)

func init() {
	initCmd.Flags().StringVar(&logic.GitDir, "gitdir", "", "specified path to git directory")
	initCmd.Flags().StringVar(&logic.WorkTree, "worktree", "", "specified path to working tree")
	initCmd.MarkFlagsRequiredTogether("gitdir", "worktree")
	rootCmd.AddCommand(initCmd)
}
