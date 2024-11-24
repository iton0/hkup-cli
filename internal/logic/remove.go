package logic

import (
	"fmt"
	"os"

	"github.com/iton0/hkup-cli/internal/git"
	"github.com/iton0/hkup-cli/internal/util"
	"github.com/spf13/cobra"
)

// Remove deletes a specified Git hook from the .hkup directory.
// It takes a single argument, which is the name of the hook to be removed.
//
// Returns error if:
//   - the .hkup directory does not exist
//   - the specified hook is not found
//   - issue deleting the file
func Remove(cmd *cobra.Command, args []string) error {
	// cannot remove if .hkup directory does not exist
	if !util.DoesDirectoryExist(util.HkupDirName) {
		return fmt.Errorf("%s directory does not exist", util.HkupDirName)
	}

	hook := args[0]

	// Validates that arg is a supported git hook
	_, err := git.GetHook(hook)
	if err != nil {
		return err
	}

	filePath := util.GetHookFilePath(hook)

	// Cannot remove if git hook does not exist in the .hkup directory
	if !util.DoesFileExist(filePath) {
		return fmt.Errorf("hook does not exist in current working directory: %s", hook)
	}

	return os.Remove(filePath)
}
