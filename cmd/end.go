package cmd

import (
	"github.com/iton0/hkup-cli/internal/logic"
	"github.com/spf13/cobra"
)

var endCmd = &cobra.Command{
	Use:   "end",
	Short: "Reset local hooksPath variable",
	Args:  cobra.NoArgs,
	RunE:  logic.End,
}

func init() {
	endCmd.Flags().BoolVar(&logic.AllFlg, "all", false, "delete local .hkup folder")
}
