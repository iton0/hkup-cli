package template

import (
	"github.com/spf13/cobra"
)

var (
	// RootCmd is the template command that will be added to the root HkUp command.
	RootCmd = &cobra.Command{
		Use:   "template",
		Short: "Reusable Git hook",
		Long:  "A template refers to a pre-configured, reusable Git hook that can be easily applied\nto a Git repository. The main goal of a template is to simplify and automate the setup\nof these hooks, making it easy to apply them consistently without having to \nwrite or configure the scripts from scratch each time.",
	}
)

func init() {
	RootCmd.AddCommand(createCmd)
	RootCmd.AddCommand(copyCmd)
	RootCmd.AddCommand(editCmd)
	RootCmd.AddCommand(removeCmd)
}
