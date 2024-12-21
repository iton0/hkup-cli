package template

import (
	"fmt"
	"os"

	"github.com/iton0/hkup-cli/internal/util"
	"github.com/spf13/cobra"
)

// Remove removes the template file from the HkUp config template directory.
// Returns error if:
//   - template does not follow naming convetion
//   - issues with removing file
func Remove(cmd *cobra.Command, args []string) error {
	templatePath := util.GetTemplateDirPath()

	if !util.DoesDirectoryExist(templatePath) {
		return fmt.Errorf("%s directory does not exist", templatePath)
	}

	templateName := args[0]

	switch file, err := doesTemplateExist(templatePath, templateName); {
	case err != nil:
		return err
	case file == "":
		return fmt.Errorf("template %s does not exist", templateName)
	default:
		return os.Remove(file)
	}
}
