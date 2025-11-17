package logic

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/iton0/hkup-cli/internal/util"
	"github.com/spf13/cobra"
)

// TemplateEdit opens specified template in the default editor for HkUp.
//
// Returns error if:
//   - template is not valid
//   - editor is not found
func TemplateEdit(_ *cobra.Command, args []string) error {
	templatePath := util.GetTemplateDirPath()

	if out, err := doesTemplateExist(templatePath, args[0]); err != nil {
		return err
	} else if out == "" {
		return fmt.Errorf("%s template does not exist", args[0])
	} else {
		return editTemplate(out)
	}
}

// editTemplate opens the template file with the default editor for HkUp.
// Returns error if issue with opening editor.
func editTemplate(path string) error {
	if editor, err := getEditor(); err != nil {
		return err
	} else {
		return util.RunCommandInTerminal(editor, path)
	}
}

// getEditor makes best effort to find default editor for HkUp.
// Returns editor name if found and error if issue with searching for editor.
func getEditor() (string, error) {
	if editor, err := util.GetINIValue("editor"); err != nil {
		return "", err
	} else if editor != "" {
		return editor, nil
	}

	out, err := exec.Command("git", "config", "--global", "core.editor").CombinedOutput()
	if len(strings.TrimSpace(string(out))) != 0 {
		if err != nil {
			return "", fmt.Errorf("issue checking global git editor configuration")
		} else {
			return string(out[:(len(out) - 1)]), nil
		}
	}

	if editor, exist := os.LookupEnv("EDITOR"); exist && editor != "" {
		return editor, nil
	}

	return "", fmt.Errorf("failed to find an editor")
}
