package logic

import (
	"os"
	"os/exec"
	"strings"

	"github.com/iton0/hkup-cli/internal/util"
	"github.com/spf13/cobra"
)

// Root wraps git-related clone commands for easier initialization of HkUp.
//
// Returns error if issue with cloning repository or initializing HkUp.
func Root(cmd *cobra.Command, args []string) error {
	// Tries to run git command in the terminal
	err := util.RunCommandInTerminal(args[0], args[1:]...)
	if err != nil {
		return err
	}

	// Tries to cd into the created directory
	if err = cdLogic(args[len(args)-2], args[len(args)-1]); err != nil {
		return err
	}

	// Tries to initialize HkUp
	return Init(cmd, nil)
}

// cdLogic implements the HkUp wrapper logic around git-related clone command.
// Returns error if issue with changing directory
func cdLogic(secondLast, dir string) error {
	// Using the regular 'git' command
	if strings.HasSuffix(dir, ".git") { // bare
		start := strings.LastIndex(dir, "/") + 1
		dir = dir[start:]

		// If the repo is not bare then updates the directory name ie regular clone
		if isBare, err := isBareRepo(dir); err != nil || !isBare {
			end := strings.LastIndex(dir, ".git")
			dir = dir[start:end]
		}
	} else if strings.Count(secondLast, "/") != 1 { // When using Github CLI
		start := strings.LastIndex(dir, "/") + 1
		dir = dir[start:]
	}

	// Either successful or returns error if issue with changing directory
	return os.Chdir(dir)
}

// isBareRepo reports if given directory (dir) is a bare git repository.
// Additionally, returns error if the given directory is not a git repository
// at all.
func isBareRepo(dir string) (bool, error) {
	// Checks if current working directory is git bare repo
	out, err := exec.Command("git", "-C", dir, "rev-parse", "--is-bare-repository").Output()
	if err != nil {
		return false, err
	}

	return strings.TrimSpace(string(out)) == "true", nil
}
