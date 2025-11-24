package cmd

import (
	"testing"

	"github.com/iton0/hkup-cli/v2/internal/testutil"
)

func TestRootCmd(t *testing.T) {
	tests := []testutil.Test{
		{
			Name:     "Valid Flag",
			Args:     []string{"-v"},
			Expected: "hkup version dev",
			Error:    false,
		},
	}

	testutil.TestCmd(t, testutil.SetupFunc(func() error {
		return nil
	}), rootCmd, tests)
}
