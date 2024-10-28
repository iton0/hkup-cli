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
	rootCmd.AddCommand(initCmd)
}
