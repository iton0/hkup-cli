package config

import (
	"github.com/iton0/hkup-cli/internal/util"
	"github.com/spf13/cobra"
)

// Get prints out the value of a specified configuration setting. Returns error
// if issue with getting the value.
func Get(cmd *cobra.Command, args []string) error {
	out, err := util.GetINIValue(args[0])
	if err != nil {
		return err
	}

	cmd.Println(out)
	return nil
}
