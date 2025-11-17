package cmd

import "github.com/spf13/cobra"

var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Reusable Git hook",
	Long: `A template refers to a pre-configured, reusable Git hook that can be easily applied
to a Git repository. The main goal of a template is to simplify and automate the setup
of these hooks, making it easy to apply them consistently without having to
write or configure the scripts from scratch each time.`,
}

func init() {
	rootCmd.AddCommand(templateCmd)
}

// FIXME: how to properly test the template subcommands without user interaction
// and changing preexisting user-defined templates
