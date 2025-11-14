package logic

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/iton0/hkup-cli/internal/util"
	"github.com/spf13/cobra"
)

// Status checks if hkup is set for the current working directory.
// Returns error if issue with getting hooksPath via git command.
func Status(cmd *cobra.Command, args []string) error {
	if !util.IsGitDirectory(".") { // Current working directory is not a git repository at all
		return fmt.Errorf("current working directory is not a git directory.\nNeed to initialize git.\n")
	}

	out, err := exec.Command("git", "config", "--local", "core.hooksPath").CombinedOutput()
	if len(strings.TrimSpace(string(out))) != 0 && err != nil {
		return fmt.Errorf("issue checking hkup status")
	} else if strings.TrimSpace(string(out)) != util.HkupDirName {
		cmd.Println("inactive")
	} else {
		cmd.Println("active")
	}

	return nil
}
