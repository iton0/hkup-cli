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
	root := args[0]

	// Tries to run git command in the terminal
	err := util.RunCommandInTerminal(root, args[1:]...)
	if err != nil {
		return err
	}

	secondLast := args[len(args)-2]
	dir := args[len(args)-1]
	var isBare bool

	// Checks if the cloned repository is bare
	for _, v := range args[1:] {
		if v == "--bare" {
			isBare = true
		}
	}

	if err = cdLogic(root, secondLast, dir, isBare); err != nil {
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
// Returns error if issue with changing directory.
func cdLogic(root, secondLast, dir string, isBare bool) error {
	// Gets the remote repository name if no custom clone name is used
	// ex). Not Custom: git clone <url>
	//      Custom: git clone <url> foo
	switch root {
	case "git", "/usr/bin/git":
		if isBare && strings.HasSuffix(dir, ".git") {
			start := strings.LastIndex(dir, "/") + 1
			dir = dir[start:]
		} else if strings.HasSuffix(dir, ".git") {
			start := strings.LastIndex(dir, "/") + 1
			end := strings.LastIndex(dir, ".git")
			dir = dir[start:end]
		}
	case "gh", "usr/bin/gh":
		if strings.Count(secondLast, "/") != 1 {
			start := strings.LastIndex(dir, "/") + 1
			dir = dir[start:]
		}
	}

	// Either successful or returns error if issue with changing directory
	return os.Chdir(dir)
}
