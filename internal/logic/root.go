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
	if err := util.RunCommandInTerminal(args[0], args[1:]...); err != nil {
		return err
	}

	possibleRepoUrl := args[len(args)-2]
	possibleCustomDir := args[len(args)-1]

	// Loop through the args to see if "--" is used.
	// Gets the two previous args and sets them to the repoUrl and
	// possibleCustomDir variables, respectively.
	for i, v := range args {
		if v == "--" {
			possibleRepoUrl = args[i-2]
			possibleCustomDir = args[i-1]
			break
		}
	}

	// Tries to cd into the created directory
	if err := cdLogic(possibleRepoUrl, possibleCustomDir); err != nil {
		return err
	}

	// Tries to initialize HkUp
	return Init(cmd, nil)
}

// cdLogic implements the HkUp wrapper logic around cloning for both 'git' and
// 'gh' command.
// Returns error if issue with changing directory
func cdLogic(possibleRepoUrl, possibleCustomDir string) error {
	// No custom directory name provided when using git command
	usedDefaultGit := strings.HasSuffix(possibleCustomDir, ".git")

	// No custom directory name provided when using gh command
	usedDefaultGh := strings.Count(possibleRepoUrl, "/") != 1 && !usedDefaultGit

	// Starting index of the remote repo name
	start := strings.LastIndex(possibleCustomDir, "/") + 1

	// Checks if user did not provide a custom directory name
	if usedDefaultGit { // git command
		// If the repo is bare then just take the remote repo name
		if isBare, _ := isBareRepo(possibleCustomDir[start:]); isBare {
			possibleCustomDir = possibleCustomDir[start:]
		} else { // Repo is not bare ie regular clone
			end := strings.LastIndex(possibleCustomDir, ".git")
			possibleCustomDir = possibleCustomDir[start:end]
		}
	} else if usedDefaultGh { // gh command
		// If the repo is bare then just take the remote repo name
		if isBare, _ := isBareRepo(possibleCustomDir[start:] + ".git"); isBare {
			possibleCustomDir = possibleCustomDir[start:] + ".git"
		} else { // Repo is not bare ie regular clone
			possibleCustomDir = possibleCustomDir[start:]
		}
	}

	// Either successful or returns error if issue with changing directory
	return os.Chdir(possibleCustomDir)
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
