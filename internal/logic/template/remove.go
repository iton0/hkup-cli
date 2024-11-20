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

	// Cannot remove template if HkUp template config path does not exist
	if !util.DoesDirectoryExist(templatePath) {
		return fmt.Errorf("%s directory does not exist.", templatePath)
	}

	templateName := args[0]

	// Checks for template existence in HkUp template config directory
	switch file, err := doesTemplateExist(templatePath, templateName); {
	case err != nil:
		return err
	case file == "": // Specified template does not exist
		return fmt.Errorf("not valid arg \"%s\" for \"hkup template remove\"", templateName)
	default: // Template exists and will try to remove
		return os.Remove(file) // Either success and returns nil or returns error
	}
}
