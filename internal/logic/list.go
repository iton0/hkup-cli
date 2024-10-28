package logic

import (
	"github.com/iton0/hkup-cli/internal/git"
	"github.com/iton0/hkup-cli/internal/util"
	"github.com/spf13/cobra"
)

// List displays a list of available Git hooks or supported languages based on the provided argument.
// It takes a single argument, which determines whether to list hooks or languages.
//
// Returns:
// - error: Returns an error if the argument is invalid; otherwise, it returns nil.
func List(cmd *cobra.Command, args []string) error {
	arg := args[0]
	var output []string

	// NOTE: default case is handled by cobra framework
	switch {
	case arg == "hook":
		output = util.ConvertMapKeysToSlice(git.Hooks())
	case arg == "lang":
		output = util.ConvertMapKeysToSlice(git.SupportedLangs())
	}

	for _, key := range output {
		cmd.Printf(" %s\n", key)
	}

	return nil
}
