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

	// ForceFlg is an optional flag which will override the local hooksPath
	// variable if it is already set.
	ForceFlg bool
)

// Init sets the .hkup directory for storing Git hooks in the current repository.
//
// Returns error if:
//   - current working directory is not a regular or bare git repository
//   - issue with creating .hkup directory
//   - hooksPath is already set
//   - issue with setting the hooksPath
func Init(cmd *cobra.Command, args []string) error {
	isBare, err := isBareRepo(".")
	if err != nil { // Current working directory is not a git repository at all
		return err
	}

	gitCmd := []string{}   // Holds everything after the root git command
	var hkupDirPath string // Holds the path the .hkup directory

	if GitDirFlg != "" && WorkTreeFlg != "" {
		hkupDirPath = filepath.Join(WorkTreeFlg, util.HkupDirName)
		gitCmd = []string{"--git-dir=" + GitDirFlg, "--work-tree=" + WorkTreeFlg, "config", "--local", "core.hooksPath", hkupDirPath}
	} else {
		hkupDirPath = util.HkupDirName
		gitCmd = []string{"config", "--local", "core.hooksPath", hkupDirPath}
	}

	if !ForceFlg {
		out, err := exec.Command("git", gitCmd[:len(gitCmd)-1]...).CombinedOutput()
		if err != nil {
			return err
		} else if len(strings.TrimSpace(string(out))) != 0 {
			return fmt.Errorf("hooksPath already set to %s", out)
		}
	}

	if err := exec.Command("git", gitCmd...).Run(); err != nil {
		return err
	}

	absPath, err := filepath.Abs(hkupDirPath)
	if err != nil {
		return err
	}
	if !util.DoesDirectoryExist(util.HkupDirName) && !isBare {
		if err := util.CreateDirectory(util.HkupDirName); err != nil {
			return err
		}

		cmd.Printf("Initialized hkup directory at %s\n", absPath)
	} else {
		cmd.Printf("Reinitialized hkup directory at %s\n", absPath)
	}

	return nil
}
