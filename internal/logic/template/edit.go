package template

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/iton0/hkup-cli/internal/util"
	"github.com/spf13/cobra"
)

// Edit opens specified template in the default editor for HkUp.
//
// Returns error if:
//   - template is not valid
//   - editor is not found
func Edit(cmd *cobra.Command, args []string) error {
	templatePath := util.GetTemplateDirPath()

	// output (without error) will either give path to template or empty string
	out, err := doesTemplateExist(templatePath, args[0])
	switch {
	case err != nil:
		return err
	case out == "":
		return fmt.Errorf("%s template does not exist", args[0])
	default:
		return editTemplate(out)
	}
}

// editTemplate opens the template file with the default editor for HkUp.
// Returns error if issue with opening editor.
func editTemplate(path string) error {
	editor, err := getEditor()
	if err != nil {
		return err
	}

	// Run command to open template file with editor
	return util.RunCommandInTerminal(editor, path) // Either success or return error
}

// getEditor makes best effort to find default editor for HkUp.
// Returns editor name if found and error if issue with searching for editor.
func getEditor() (string, error) {
	// Check the HkUp config file
	editor, err := util.GetINIValue("editor")
	if err != nil {
		return "", err
	} else if editor != "" {
		return editor, nil
	}

	// Check in global gitconfig file
	if out, err := exec.Command("git", "config", "--global", "core.editor").CombinedOutput(); err != nil {
		return "", err
	} else if len(out) != 0 {
		// The out has a newline character at the end so take elements up until the
		// "\" of the "\n"
		return string(out[0:(len(out) - 1)]), nil // Converts byte slice into string
	}

	// Check for EDITOR var
	if editor, exist := os.LookupEnv("EDITOR"); exist && editor != "" {
		return editor, nil
	}

	return "", fmt.Errorf("failed to find an editor")
}
