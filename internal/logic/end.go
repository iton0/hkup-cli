package logic

import (
	"fmt"
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
func End(cmd *cobra.Command, args []string) error {
	if AllFlg {
		if err := os.RemoveAll(util.HkupDirName); err != nil {
			return fmt.Errorf("issue removing %s", util.HkupDirName)
		}
		cmd.Println("Removed .hkup directory and contents")
	}

	err := exec.Command("git", "config", "--local", "core.hooksPath", "").Run()
	if err != nil {
		return fmt.Errorf("issue unsetting local hooksPath variable")
	}

	cmd.Println("hkup unset")
	return nil
}
