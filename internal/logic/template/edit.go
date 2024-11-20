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
	editor, err := util.GetEditor()
	if err != nil {
		return err
	}

	// Create the command to open the editor with the template file
	cmd := exec.Command(editor, path)

	// This allows the editor to be opened in the same terminal
	// Source: https://stackoverflow.com/questions/12088138/trying-to-launch-an-external-editor-from-within-a-go-program#12089980
	// NOTE: This only applies to terminal-based editors such as vim, nvim, etc.
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Starts the editor
	err = cmd.Start()
	if err != nil {
		return err
	}

	// Waits for the user to finish editing
	return cmd.Wait() // Either success and returns nil or returns error if issue
}
