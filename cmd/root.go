package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	// NOTE: Updated to the latest release version at build time of the binaries
	//       via .github/workflows/release-please.yml.
	// 		 Defaults to "dev" when developing/building
	version = "dev"

	rootCmd = &cobra.Command{
		Use:     "hkup",
		Short:   "hkup CLI",
		Long:    `hkup is a management tool for git hooks`,
		Args:    cobra.MinimumNArgs(1),
		Version: version,
	}
)

func init() {
}

// Execute serves as a wrapper for the Cobra API's Execute function, allowing it
// to be called from the [github.com/iton0/hkup-cli] package.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
