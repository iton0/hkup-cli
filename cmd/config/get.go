package config

import (
	"github.com/iton0/hkup-cli/internal/logic/config"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get <config-setting>",
	Short: "Get a HkUp config setting",
	Args:  cobra.ExactArgs(1),
	RunE:  config.Get,
}

func init() {}
