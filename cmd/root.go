package cmd

import (
	"github.com/iton0/hkup-cli/cmd/config"
	"github.com/iton0/hkup-cli/cmd/template"
	"github.com/iton0/hkup-cli/internal/logic"
	"github.com/spf13/cobra"
)

var (
	// NOTE: Updated to the latest release version at build time of the binaries
	//       via .github/workflows/release-please.yml.
	// 		 Defaults to "dev" when developing/building
	version = "dev"

	rootCmd = &cobra.Command{
		Use:     "hkup [-- <git/gh clone command>]",
		Short:   "hkup CLI",
		Long:    `hkup is a management tool for git hooks`,
		Args:    cobra.MinimumNArgs(1),
		Version: version,
		RunE:    logic.Root,
	}
)

func init() {
	rootCmd.AddCommand(endCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(template.RootCmd)
	rootCmd.AddCommand(config.RootCmd)
	rootCmd.AddCommand(docCmd)
	rootCmd.AddCommand(listCmd)
}

// Execute serves as a wrapper for the Cobra API's Execute function, allowing it
// to be called from the [github.com/iton0/hkup-cli] package.
func Execute() {
	rootCmd.Execute()
}
