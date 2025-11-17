package cmd

import (
	"testing"

	"github.com/iton0/hkup-cli/internal/testutil"
)

func TestVersionCmd(t *testing.T) {
	tests := []testutil.Test{
		{
			Name:     "Valid Arg",
			Args:     []string{"version"},
			Expected: "hkup version dev",
			Error:    false,
		},
	}

	testutil.TestCmd(t, testutil.SetupFunc(func() error {
		return nil
	}), rootCmd, tests)
}
