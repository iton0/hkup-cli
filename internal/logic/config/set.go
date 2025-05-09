package config

import (
	"github.com/iton0/hkup-cli/internal/util"
	"github.com/spf13/cobra"
)

// Set updates a configuration setting with a new value. Returns error if issue
// with settings the configuration setting.
func Set(_ *cobra.Command, args []string) error {
	return util.SetINIValue(args[0], args[1])
}
