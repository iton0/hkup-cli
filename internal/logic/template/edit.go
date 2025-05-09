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
func Edit(_ *cobra.Command, args []string) error {
	templatePath := util.GetTemplateDirPath()

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

	return util.RunCommandInTerminal(editor, path)
}

// getEditor makes best effort to find default editor for HkUp.
// Returns editor name if found and error if issue with searching for editor.
func getEditor() (string, error) {
	editor, err := util.GetINIValue("editor")
	if err != nil {
		return "", err
	} else if editor != "" {
		return editor, nil
	}

	if out, err := exec.Command("git", "config", "--global", "core.editor").CombinedOutput(); err != nil && len(out) != 0 {
		return "", err
	} else if len(out) != 0 {
		return string(out[0:(len(out) - 1)]), nil
	}

	if editor, exist := os.LookupEnv("EDITOR"); exist && editor != "" {
		return editor, nil
	}

	return "", fmt.Errorf("failed to find an editor")
}
