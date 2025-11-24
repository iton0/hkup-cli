package logic

import (
	"github.com/iton0/hkup-cli/v2/internal/util"
	"github.com/spf13/cobra"
)

// ConfigSet updates a configuration setting with a new value. Returns error if issue
// with settings the configuration setting.
func ConfigSet(_ *cobra.Command, args []string) error {
	return util.SetINIValue(args[0], args[1])
}
