package cmd

import (
	"github.com/iton0/hkup-cli/internal/git"
	"github.com/iton0/hkup-cli/internal/logic"
	"github.com/iton0/hkup-cli/internal/util"
	"github.com/spf13/cobra"
)

var templateCreateCmd = &cobra.Command{
	Use:       "create [<hook-name>]",
	Short:     "Create a git hook template",
	ValidArgs: util.ConvertMapKeysToSlice(git.Hooks()),
	Args:      cobra.MatchAll(cobra.MaximumNArgs(1), cobra.OnlyValidArgs),
	RunE:      logic.TemplateCreate,
}

func init() {
	templateCreateCmd.Flags().StringVar(&logic.TemplateLangFlg, "lang", "", "supported languages for git hooks")
	templateCreateCmd.Flags().StringVar(&logic.TemplateNameFlg, "name", "", "specified name for git hook template")
	templateCreateCmd.Flags().BoolVar(&logic.TemplateCwdFlg, "cwd", false, "use hook from current working directory")
	templateCreateCmd.Flags().BoolVar(&logic.TemplateCopyFlg, "copy", false, "copy to current working directory")
	templateCreateCmd.Flags().BoolVar(&logic.TemplateEditFlg, "edit", false, "open template in editor")
	templateCreateCmd.MarkFlagsMutuallyExclusive("cwd", "lang")
	templateCreateCmd.MarkFlagsMutuallyExclusive("cwd", "copy")
	templateCmd.AddCommand(templateCreateCmd)
}
