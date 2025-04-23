package logic

import (
	"os"
	"os/exec"

	"github.com/iton0/hkup-cli/internal/util"
	"github.com/spf13/cobra"
)

// AllFlg is a optional flag that instructs to delete the local .hkup folder
var AllFlg bool

// End resets the local hooksPath variable
// Returns error if issue:
// - deleting .hkup folder (if --all flag used)
// - resetting the hooksPath
func End(_ *cobra.Command, args []string) error {
	if AllFlg {
		if err := os.RemoveAll(util.HkupDirName); err != nil {
			return err
		}
	}

	return exec.Command("git", "config", "--local", "core.hooksPath", "").Run()
}
