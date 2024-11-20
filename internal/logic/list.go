package logic

import (
	"github.com/iton0/hkup-cli/internal/git"
	"github.com/iton0/hkup-cli/internal/util"
	"github.com/spf13/cobra"
)

// List displays a list of available Git hooks or supported languages based on
// the provided argument.
//
// Returns error if the argument is invalid.
func List(cmd *cobra.Command, args []string) error {
	arg := args[0]
	out := []string{}

	// NOTE: Default case is handled by cobra framework
	switch {
	case arg == "hook":
		out = util.ConvertMapKeysToSlice(git.Hooks())
	case arg == "lang":
		out = util.ConvertMapKeysToSlice(git.SupportedLangs())
	}

	for _, key := range out {
		cmd.Printf(" %s\n", key)
	}

	return nil
}
