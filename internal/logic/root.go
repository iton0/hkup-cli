package logic

import (
	"os"
	"os/exec"
	"strings"

	"github.com/iton0/hkup-cli/internal/util"
	"github.com/spf13/cobra"
)

// Root wraps git-related clone commands for easier initialization of HkUp.
// Returns error if issue with cloning repository or initializing HkUp.
func Root(cmd *cobra.Command, args []string) error {
	if err := util.RunCommandInTerminal(args[0], args[1:]...); err != nil {
		return err
	}

	possibleRepoUrl := args[len(args)-2]
	possibleCustomDir := args[len(args)-1]

	// INFO: Since gh command can use the syntax of:
	//   gh repo clone <url> [<customdir>] -- <gitflags>
	// Need to parse for "--" and update possible repo url and custom directory
	for i, v := range args {
		if v == "--" {
			possibleRepoUrl = args[i-2]
			possibleCustomDir = args[i-1]
			break
		}
	}

	if err := cdLogic(possibleRepoUrl, possibleCustomDir); err != nil {
		return err
	}

	return Init(cmd, nil)
}

// cdLogic implements the HkUp wrapper logic around cloning for both 'git' and
// 'gh' command.
// Returns error if issue with changing directory.
func cdLogic(possibleRepoUrl, possibleCustomDir string) error {
	// No custom directory name provided when using git command
	usedDefaultGit := strings.HasSuffix(possibleCustomDir, ".git")

	// No custom directory name provided when using gh command
	usedDefaultGh := strings.Count(possibleRepoUrl, "/") != 1 && !usedDefaultGit

	// Starting index of the remote repo name
	start := strings.LastIndex(possibleCustomDir, "/") + 1

	createdDir := possibleCustomDir // Holds the name of the cloned directory

	if usedDefaultGit {
		if isBare, _ := isBareRepo(possibleCustomDir[start:]); isBare {
			createdDir = possibleCustomDir[start:]
		} else {
			end := strings.LastIndex(possibleCustomDir, ".git")
			createdDir = possibleCustomDir[start:end]
		}
	} else if usedDefaultGh {
		if isBare, _ := isBareRepo(possibleCustomDir[start:] + ".git"); isBare {
			createdDir = possibleCustomDir[start:] + ".git"
		} else {
			createdDir = possibleCustomDir[start:]
		}
	}

	return os.Chdir(createdDir)
}

// isBareRepo reports if given directory (dir) is a bare git repository.
// Additionally, returns error if the given directory is not a git repository
// at all.
func isBareRepo(dir string) (bool, error) {
	out, err := exec.Command("git", "-C", dir, "rev-parse", "--is-bare-repository").Output()
	if err != nil {
		return false, err
	}

	return strings.TrimSpace(string(out)) == "true", nil
}
