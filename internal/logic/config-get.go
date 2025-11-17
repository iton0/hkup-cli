package logic

import (
	"github.com/iton0/hkup-cli/internal/util"
	"github.com/spf13/cobra"
)

// ConfigGet prints out the value of a specified configuration setting. Returns error
// if issue with getting the value.
func ConfigGet(cmd *cobra.Command, args []string) error {
	if out, err := util.GetINIValue(args[0]); err != nil {
		return err
	} else {
		cmd.Println(out)
		return nil
	}
}
