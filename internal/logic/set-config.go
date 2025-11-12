package logic

import (
	"github.com/iton0/hkup-cli/internal/util"
	"github.com/spf13/cobra"
)

// SetConfig updates a configuration setting with a new value. Returns error if issue
// with settings the configuration setting.
func SetConfig(_ *cobra.Command, args []string) error {
	return util.SetINIValue(args[0], args[1])
}
