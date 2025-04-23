package logic

import (
	"os/exec"
	"strings"

	"github.com/iton0/hkup-cli/internal/util"
	"github.com/spf13/cobra"
)

// Status checks if hkup is set for the current working directory.
// Returns error if issue with getting hooksPath via git command.
func Status(cmd *cobra.Command, args []string) error {
	_, err := isBareRepo(".")
	if err != nil { // Current working directory is not a git repository at all
		cmd.Printf("Current working directory is not a git repository\n")
		return nil
	}

	if out, err := exec.Command("git", "config", "--local", "core.hooksPath").CombinedOutput(); err != nil && len(out) != 0 {
		return err
	} else if strings.TrimSpace(string(out)) != util.HkupDirName {
		cmd.Printf("inactive\n")
		return nil
	}

	cmd.Printf("active\n")
	return nil
}
