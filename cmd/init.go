package cmd

import (
	"github.com/iton0/hkup-cli/internal/logic"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize hkup",
	Long:  "Create an empty hkup directory or reinitialize an existing one",
	Args:  cobra.NoArgs,
	RunE:  logic.Init,
}

func init() {
	initCmd.Flags().BoolVarP(&logic.ForceFlg, "force", "f", false, "override local hooksPath variable")
	initCmd.Flags().StringVar(&logic.GitDirFlg, "gitdir", "", "specified path to git directory")
	initCmd.Flags().StringVar(&logic.WorkTreeFlg, "worktree", "", "specified path to working tree")
	initCmd.MarkFlagsRequiredTogether("gitdir", "worktree")
	rootCmd.AddCommand(initCmd)
}
