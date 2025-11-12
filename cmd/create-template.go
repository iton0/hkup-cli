package cmd

import (
	"github.com/iton0/hkup-cli/internal/git"
	"github.com/iton0/hkup-cli/internal/logic"
	"github.com/iton0/hkup-cli/internal/util"
	"github.com/spf13/cobra"
)

var createTemplateCmd = &cobra.Command{
	Use:       "create [<hook-name>]",
	Short:     "Create a git hook template",
	ValidArgs: util.ConvertMapKeysToSlice(git.Hooks()),
	Args:      cobra.MatchAll(cobra.MaximumNArgs(1), cobra.OnlyValidArgs),
	RunE:      logic.CreateTemplate,
}

func init() {
	createTemplateCmd.Flags().StringVar(&logic.TemplateLangFlg, "lang", "", "supported languages for git hooks")
	createTemplateCmd.Flags().StringVar(&logic.TemplateNameFlg, "name", "", "specified name for git hook template")
	createTemplateCmd.Flags().BoolVar(&logic.TemplateCwdFlg, "cwd", false, "use hook from current working directory")
	createTemplateCmd.Flags().BoolVar(&logic.TemplateCopyFlg, "copy", false, "copy to current working directory")
	createTemplateCmd.Flags().BoolVar(&logic.TemplateEditFlg, "edit", false, "open template in editor")
	createTemplateCmd.MarkFlagsMutuallyExclusive("cwd", "lang")
	createTemplateCmd.MarkFlagsMutuallyExclusive("cwd", "copy")
	templateCmd.AddCommand(createTemplateCmd)
}
