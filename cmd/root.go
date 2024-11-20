package cmd

import (
	"github.com/iton0/hkup-cli/cmd/config"
	"github.com/iton0/hkup-cli/cmd/template"
	"github.com/spf13/cobra"
)

var (
	// version holds the centralized version of HkUp.
	// It is updated to the latest release version at build time of the binaries.
	//
	// INFO: look at the .github/workflows/release-please.yml to view how version
	//       is updated.
	version = "dev"

	rootCmd = &cobra.Command{
		Use:     "hkup",
		Short:   "hkup CLI",
		Long:    `hkup is a management tool for git hooks`,
		Version: version,
	}
)

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(template.RootCmd)
	rootCmd.AddCommand(config.RootCmd)
	rootCmd.AddCommand(docCmd)
	rootCmd.AddCommand(listCmd)
}

// Execute serves as a wrapper for the Cobra API's Execute function,
// allowing it to be called from the [github.com/iton0/hkup-cli] package.
func Execute() {
	rootCmd.Execute()
}
