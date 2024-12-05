package logic

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

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
//   - current working directory is not a git repository/worktree
//   - issue with creating .hkup directory
//   - hooksPath is already set
//   - issue with setting the hooksPath
func Init(cmd *cobra.Command, args []string) error {
	// Checks if current working directory is git worktree
	out, err := exec.Command("git", "-C", ".", "rev-parse", "--is-inside-work-tree").Output()
	if err != nil {
		return err
	}

	// Current working directory is not a worktree so returns error
	result := strings.TrimSpace(string(out))
	// If the WorkTree flag is used cwd does not need to be a git worktree
	if result == "false" && WorkTreeFlg == "" {
		return fmt.Errorf("must run \"hkup init\" inside a worktree")
	}

	var gitCmd []string    // Holds everything after the root git command
	var hkupDirPath string // Holds the path the .hkup directory

	// If both flags are set, configure core.hooksPath with their values.
	// Otherwise, use the default hooks directory (util.HkupDirName).
	if GitDirFlg != "" && WorkTreeFlg != "" {
		hkupDirPath = filepath.Join(WorkTreeFlg, util.HkupDirName)
		gitCmd = []string{"--git-dir=" + GitDirFlg, "--work-tree=" + WorkTreeFlg, "config", "--local", "core.hooksPath", hkupDirPath}
	} else {
		hkupDirPath = util.HkupDirName
		gitCmd = []string{"config", "--local", "core.hooksPath", hkupDirPath}
	}

	// Tries to create the .hkup directory if it does not exist
	if !util.DoesDirectoryExist(hkupDirPath) {
		if err = util.CreateDirectory(hkupDirPath); err != nil {
			return err
		}

		cmd.Printf("Initialized hkup directory at %s\n", hkupDirPath)
	}

	// Does not override the hooksPath variable if already set
	if out, _ := exec.Command("git", gitCmd[:len(gitCmd)-1]...).CombinedOutput(); len(out) != 0 {
		return fmt.Errorf("hooksPath already set to %s", out)
	}

	return exec.Command("git", gitCmd...).Run()
}
