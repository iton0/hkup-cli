package cmd

import (
	"github.com/iton0/hkup-cli/internal/git"
	"github.com/iton0/hkup-cli/internal/logic/template"
	"github.com/iton0/hkup-cli/internal/util"
	"github.com/spf13/cobra"
)

var createTemplateCmd = &cobra.Command{
	Use:       "create [<hook-name>]",
	Short:     "Create a git hook template",
	ValidArgs: util.ConvertMapKeysToSlice(git.Hooks()),
	Args:      cobra.MatchAll(cobra.MaximumNArgs(1), cobra.OnlyValidArgs),
	RunE:      template.Create,
}

func init() {
	createTemplateCmd.Flags().StringVar(&template.TemplateLangFlg, "lang", "", "supported languages for git hooks")
	createTemplateCmd.Flags().StringVar(&template.TemplateNameFlg, "name", "", "specified name for git hook template")
	createTemplateCmd.Flags().BoolVar(&template.TemplateCwdFlg, "cwd", false, "use hook from current working directory")
	createTemplateCmd.Flags().BoolVar(&template.TemplateCopyFlg, "copy", false, "copy to current working directory")
	createTemplateCmd.Flags().BoolVar(&template.TemplateEditFlg, "edit", false, "open template in editor")
	createTemplateCmd.MarkFlagsMutuallyExclusive("cwd", "lang")
	createTemplateCmd.MarkFlagsMutuallyExclusive("cwd", "copy")
	templateCmd.AddCommand(createTemplateCmd)
}
