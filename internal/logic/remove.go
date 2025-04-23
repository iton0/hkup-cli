package logic

import (
	"os"

	"github.com/iton0/hkup-cli/internal/util"
	"github.com/spf13/cobra"
)

// Remove deletes a specified Git hook from the .hkup directory.
// It takes a single argument, which is the name of the hook to be removed.
// Returns error if issue deleting the file.
func Remove(_ *cobra.Command, args []string) error {
	return os.Remove(util.GetHookFilePath(args[0]))
}
