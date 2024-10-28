package logic

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/iton0/hkup-cli/internal/util"
	"github.com/spf13/cobra"
)

// Remove deletes a specified Git hook from the hkup folder.
// It takes a single argument, which is the name of the hook to be removed.
//
// Returns:
//   - error: Returns an error if the hkup folder does not exist, if the specified hook is not found,
//     or if there is an issue deleting the file; otherwise, it returns nil.
func Remove(cmd *cobra.Command, args []string) error {
	hook := args[0]

	if !util.DoesDirectoryExist(FullPath) {
		return fmt.Errorf("failed running \"hkup remove\"\n%s folder does not exist", FullPath)
	}

	filePath := filepath.Join(FullPath, hook)

	if !util.DoesFileExist(filePath) {
		return fmt.Errorf("not supported hook: %s", hook)
	}

	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("failed deleting file: %w", err)
	}

	return nil
}
