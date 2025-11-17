package cmd

import (
	"testing"

	"github.com/iton0/hkup-cli/internal/testutil"
)

func TestConfigGetCmd(t *testing.T) {
	tests := []testutil.Test{
		{
			Name:     "Valid Arg for Editor",
			Args:     []string{"config", "get", "editor"},
			Expected: "",
			Error:    false,
		},
		{
			Name:     "Invalid Key",
			Args:     []string{"config", "get", "sandwich"},
			Expected: `Error: "sandwich" is not a valid key`,
			Error:    true,
		},
	}

	testutil.TestCmd(t, testutil.SetupFunc(func() error {
		return nil
	}), rootCmd, tests)
}
