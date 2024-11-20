package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "hkup",
		Short:   "hkup CLI",
		Long:    `hkup is a management tool for git hooks`,
		Args:    cobra.MinimumNArgs(1),
		Version: "0.2.1",
	}
)

func init() {}
	rootCmd.AddCommand(template.RootCmd)

// Execute serves as a wrapper for the Cobra API's Execute function,
// allowing it to be called from the main package.
func Execute() {
	rootCmd.Execute()
}
