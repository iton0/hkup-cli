package logic

import (
	"fmt"
	"os"

	"github.com/iton0/hkup-cli/internal/util"
	"github.com/spf13/cobra"
)

// RemoveTemplate removes the template file from the HkUp config template directory.
// Returns error if:
//   - template does not follow naming convetion
//   - issues with removing file
func RemoveTemplate(_ *cobra.Command, args []string) error {
	templatePath := util.GetTemplateDirPath()

	if !util.DoesDirectoryExist(templatePath) {
		return fmt.Errorf("%s directory does not exist", templatePath)
	}

	templateName := args[0]

	if file, err := doesTemplateExist(templatePath, templateName); err != nil {
		return err
	} else if file == "" {
		return fmt.Errorf("template %s does not exist", templateName)
	} else {
		return os.Remove(file)
	}
}
