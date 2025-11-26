package cmd

import "github.com/spf13/cobra"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Retrieve hkup version",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Printf("hkup version %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
