package template

import (
	"github.com/iton0/hkup-cli/internal/git"
	"github.com/iton0/hkup-cli/internal/logic/template"
	"github.com/iton0/hkup-cli/internal/util"
	"github.com/spf13/cobra"
)

var (
	createCmd = &cobra.Command{
		Use:       "create [<hook-name>]",
		Short:     "Create a git hook template",
		ValidArgs: util.ConvertMapKeysToSlice(git.Hooks()),
		Args:      cobra.MatchAll(cobra.MaximumNArgs(1), cobra.OnlyValidArgs),
		RunE:      template.Create,
	}
)

func init() {
	createCmd.Flags().StringVar(&template.TemplateLangFlg, "lang", "", "supported languages for git hooks")
	createCmd.Flags().StringVar(&template.TemplateNameFlg, "name", "", "specified name for git hook template")
	createCmd.Flags().BoolVar(&template.TemplateCwdFlg, "cwd", false, "use hook from current working directory")
	createCmd.Flags().BoolVar(&template.TemplateCopyFlg, "copy", false, "copy to current working directory")
	createCmd.Flags().BoolVar(&template.TemplateEditFlg, "edit", false, "open template in editor")
	createCmd.MarkFlagsMutuallyExclusive("cwd", "lang")
	createCmd.MarkFlagsMutuallyExclusive("cwd", "copy")
}
