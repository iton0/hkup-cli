package logic

import (
	"fmt"
	"os"

	"github.com/iton0/hkup-cli/internal/git"
	"github.com/iton0/hkup-cli/internal/util"
	"github.com/spf13/cobra"
)

// List displays a list of one of the below based on provided argument:
//   - Supported git hooks
//   - Supported languages
//   - User-defined git hook templates
//   - Git hook(s) used in the current working directory
//
// Returns error if issue with checking directories.
func List(cmd *cobra.Command, args []string) error {
	out := []string{}

	if len(args) > 0 {
		switch args[0] {
		case "template":
			out = getHookTemplates()
			if out == nil {
				return fmt.Errorf("could not read template directory")
			}
		case "hook":
			out = util.ConvertMapKeysToSlice(git.Hooks())
		case "lang":
			out = util.ConvertMapKeysToSlice(git.SupportedLangs())
		}
	} else {
		out = getCwdHooks()
		if out == nil {
			return fmt.Errorf("could not read .hkup directory")
		}
	}

	cmd.Print(formatOutput(out))
	return nil
}

// formatOutput formats the output string slice as a string that is returned
func formatOutput(out []string) string {
	var fout string

	for _, val := range out {
		fout += " " + val + "\n"
	}

	return fout
}

// getHookTemplates returns a slice of all user-defined template file names.
// If no templates are found, returns a empty slice.
func getHookTemplates() []string {
	out := []string{}

	files, err := os.ReadDir(util.GetTemplateDirPath())
	if err != nil {
		return nil
	}

	for _, file := range files {
		out = append(out, file.Name())
	}

	return out
}

// getCwdHooks returns the hooks in the current working directory.
// If no hooks are found, returns a empty string slice.
func getCwdHooks() []string {
	files, err := os.ReadDir(util.HkupDirName)
	if err != nil {
		return nil
	}

	out := []string{}

	for _, file := range files {
		out = append(out, file.Name())
	}

	return out
}
