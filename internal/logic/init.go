package logic

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/iton0/hkup-cli/internal/util"
	"github.com/spf13/cobra"
)

var (
	// FullPath defines the local repository folder name to hold Git hooks via a
	// relative path.
	// It is treated as a constant and points to the ".hkup" directory within the
	// current working directory.
	FullPath = filepath.Join(".", ".hkup")

	// gitCmd defines the terminal command to use to initialize the hkup folder.
	gitCmd = []string{"config", "--local", "core.hooksPath", FullPath}

	// GitDir is an optional flag that defines the location of the git directory.
	// Can be useful for bare repos or custom git setups where git directory and
	// working directory are not in the same location.
	GitDir string

	// WorkTree is an optional flag that defines the location of the working tree
	// of a local git repository.
	WorkTree string
)

// Init initializes the hkup folder for storing Git hooks in the current repository.
// It checks if there the git directory and worktree flags are provided.
// Else it will check if the current working directory is a Git repository, creates
// the hkup folder if it doesn't exist, and sets the Git configuration for
// `core.hooksPath` to point to the hkup folder.
// Returns an error if the current directory is not a Git repository, if the folder
// creation fails, or if there is an issue setting the Git hooks path.
//
// Returns:
// - error: Returns an error if any of the steps fail; otherwise, it returns nil.
func Init(cmd *cobra.Command, args []string) error {
	if GitDir != "" && WorkTree != "" {
		gitCmd = []string{"--git-dir=" + GitDir, "--work-tree=" + WorkTree, "config", "--local", "core.hooksPath", FullPath}
	} else if err := exec.Command("git", "-C", ".", "rev-parse", "--is-inside-work-tree").Run(); err != nil {
		return fmt.Errorf("failed to check if current working directory is git repo: %w", err)
	}

	if !util.DoesDirectoryExist(FullPath) {
		if err := util.CreateFolder(FullPath); err != nil {
			return err
		}
		cmd.Printf("Initialized hkup folder at %s\n", FullPath)
	}

	if out, _ := exec.Command("git", "config", "--local", "core.hooksPath").Output(); len(out) != 0 {
		return fmt.Errorf("hooksPath already set to %s", out)
	} else {
		if err := exec.Command("git", gitCmd...).Run(); err != nil {
			return fmt.Errorf("failed to set hooksPath: %w", err)
		}
		return nil
	}
}
