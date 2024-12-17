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

	secondLast := args[len(args)-2] // Second last arg that user provided
	dir := args[len(args)-1]        // Last arg; assume it is the directory
	var isBare bool                 // Boolean whether directory is bare

	// Tries to cd into the created directory and updates the isBare variable
	if isBare, err = cdLogic(secondLast, dir); err != nil {
		return err
	}

	// Current working directory is a bare repository and
	// prompts the user to create a worktree & cd into the worktree
	if isBare {
		// Asks for the worktree path
		worktree, err := util.UserInputPrompt("Create Worktree Path:")
		if err != nil {
			return err
		}

		// Asks for the branch to checkout worktree into
		branch, err := util.UserInputPrompt("Branch:")
		if err != nil {
			return err
		}

		// Tries to create worktree and cd into it
		err = exec.Command("git", "worktree", "add", worktree, branch).Run()
		if err != nil {
			return err
		}
		if err = os.Chdir(worktree); err != nil {
			return err
		}
	}

	// Tries to initialize HkUp
	return Init(cmd, nil)
}

// cdLogic implements the HkUp wrapper logic around git-related clone command.
// Returns:
//   - boolean of if directory is bare
//   - error if issue with changing directory
func cdLogic(secondLast, dir string) (bool, error) {
	// Gets the remote repository name if no custom clone name is used
	// ex). Not Custom: git clone <url>
	//      Custom: git clone <url> foo

	// Holds the boolean whether directory is bare or not
	var isBare bool

	// Using the regular 'git' command
	if strings.HasSuffix(dir, ".git") { // bare
		start := strings.LastIndex(dir, "/") + 1
		dir = dir[start:]

		// If the repo is not bare then updates the directory name ie regular clone
		if isBare = isBareRepo(dir); !isBare {
			end := strings.LastIndex(dir, ".git")
			dir = dir[start:end]
		}
	} else if strings.Count(secondLast, "/") != 1 { // When using Github CLI
		start := strings.LastIndex(dir, "/") + 1
		dir = dir[start:]
	}

	// Either successful or returns error if issue with changing directory
	return isBare, os.Chdir(dir)
}

// isBareRepo reports if given directory (dir) is a bare git repository.
func isBareRepo(dir string) bool {
	// Checks if current working directory is git bare repo
	out, err := exec.Command("git", "-C", dir, "rev-parse", "--is-bare-repository").Output()
	if err != nil {
		return false
	}

	// Returns error if current working directory is not a worktree
	result := strings.TrimSpace(string(out))

	return result == "true"
}
