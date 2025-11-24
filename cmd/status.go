package cmd

import (
	"github.com/iton0/hkup-cli/v2/internal/logic"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get status of hkup",
	Long:  "Shows whether hkup is initialized for the current working directory",
	Args:  cobra.NoArgs,
	RunE:  logic.Status,
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
