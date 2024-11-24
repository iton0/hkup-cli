package config

import (
	"github.com/iton0/hkup-cli/internal/util"
	"github.com/spf13/cobra"
)

// Set updates a configuration setting with a new value.
//
// Returns error if issue with settings the configuration setting.
func Set(cmd *cobra.Command, args []string) error {
	// Either setting configuration setting is successful and returns nil or
	// returns error
	return util.SetINIValue(util.GetConfigFilePath(), args[0], args[1])
}
