package cmd

import (
	"bytes"
	"testing"

	"github.com/iton0/hkup-cli/internal/testutil"
)

func TestEndCmd(t *testing.T) {
	tests := []testutil.Test{
		{
			Name:     "Valid Arg with Flag",
			Args:     []string{"end", "--all"},
			Expected: "Removed .hkup directory and contents\nhkup unset",
			Error:    false,
		},
	}

	testutil.TestCmd(t, testutil.SetupFunc(func() error {
		outBuf := new(bytes.Buffer)
		errBuf := new(bytes.Buffer)

		rootCmd.SetOut(outBuf)
		rootCmd.SetErr(errBuf)
		rootCmd.SetArgs([]string{"init"})

		return rootCmd.Execute()
	}), rootCmd, tests)
}
