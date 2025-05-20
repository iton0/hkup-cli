package config

import (
	"github.com/iton0/hkup-cli/internal/util"
	"github.com/spf13/cobra"
)

// Get prints out the value of a specified configuration setting. Returns error
// if issue with getting the value.
func Get(cmd *cobra.Command, args []string) error {
	if out, err := util.GetINIValue(args[0]); err != nil {
		return err
	} else {
		cmd.Println(out)
		return nil
	}
}
