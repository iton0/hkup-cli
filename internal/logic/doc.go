package logic

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/iton0/hkup-cli/internal/git"
	"github.com/spf13/cobra"
)

// Doc opens the documentation for a specified Git hook in the default web browser.
// The command takes a single argument, which is the key (name) of the hook.
// It constructs the URL for the documentation based on the hook name and attempts
// to open it using the appropriate command for the operating system.
//
// Returns:
//   - error if the hook key is invalid, if the platform is unsupported, or if
//     there is an issue starting the command.
func Doc(cmd *cobra.Command, args []string) error {
	// Full url path for the specified git hook
	url := "https://git-scm.com/docs/githooks#" + git.GetHookUrl(args[0])

	var termCmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		termCmd = exec.Command("xdg-open", url)
	case "darwin":
		termCmd = exec.Command("open", url)
	case "windows":
		termCmd = exec.Command("explorer", url)
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}

	if err := termCmd.Start(); err != nil {
		return err
	}

	// Must be called after successfully starting terminal command above
	return termCmd.Wait() // Returns error if command fails
}
