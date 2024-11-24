package logic

import (
	"fmt"
	"os/exec"

	"github.com/iton0/hkup-cli/internal/util"
	"github.com/spf13/cobra"
)

var (
	// GitDirFlg is an optional flag that defines the location of the git directory.
	// Can be useful for bare repos or custom git setups where git directory and
	// working directory are not in the same location.
	GitDirFlg string

	// WorkTreeFlg is an optional flag that defines the location of the working tree
	// of a local git repository.
	WorkTreeFlg string
)

// Init sets the .hkup directory for storing Git hooks in the current repository.
//
// Returns error if:
//   - current working directory is not a git repo
//   - issue with creating .hkup directory
//   - hooksPath is already set
//   - issue with setting the hooksPath
func Init(cmd *cobra.Command, args []string) error {
	// Only runs if current working directory is git repo
	err := exec.Command("git", "-C", ".", "rev-parse", "--is-inside-work-tree").Run()
	if err != nil {
		return err
	}

	// Tries to create the .hkup directory if it does not exist
	if !util.DoesDirectoryExist(util.HkupDirName) {
		if err = util.CreateDirectory(util.HkupDirName); err != nil {
			return err
		}

		cmd.Printf("Initialized hkup directory at %s\n", util.HkupDirName)
	}

	// Does not override the hooksPath variable if already set
	if out, _ := exec.Command("git", "config", "--local", "core.hooksPath").CombinedOutput(); len(out) != 0 {
		return fmt.Errorf("hooksPath already set to %s", out)
	}

	// Holds everything after the base 'git' in the command
	var gitCmd []string

	if GitDirFlg != "" && WorkTreeFlg != "" {
		gitCmd = []string{"--git-dir=" + GitDirFlg, "--work-tree=" + WorkTreeFlg, "config", "--local", "core.hooksPath", util.HkupDirName}
	} else {
		gitCmd = []string{"config", "--local", "core.hooksPath", util.HkupDirName}
	}

	return exec.Command("git", gitCmd...).Run()
}
