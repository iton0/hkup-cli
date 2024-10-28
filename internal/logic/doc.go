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
// It constructs the URL for the documentation based on the hook name and attempts to open it using the appropriate command for the operating system.
//
// Returns:
// - error: Returns an error if the hook key is invalid, if the platform is unsupported, or if there is an issue starting the command; otherwise, it returns nil.
func Doc(cmd *cobra.Command, args []string) error {
	key := args[0]
	var termCmd *exec.Cmd
	var url string

	if hook, err := git.GetHook(key); err != nil {
		return err
	} else {
		url = git.HookDocSite + hook
	}

	switch runtime.GOOS {
	case "linux":
		termCmd = exec.Command("xdg-open", url)
	case "darwin":
		termCmd = exec.Command("open", url)
	case "windows":
		termCmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}

	termCmd.Start()

	return termCmd.Wait()
}
