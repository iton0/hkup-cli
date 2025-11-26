package cmd

import (
	"bytes"
	"testing"

	"github.com/iton0/hkup-cli/v2/internal/testutil"
)

func TestRemoveCmd(t *testing.T) {
	tests := []testutil.Test{
		{
			Name:     "Valid Arg",
			Args:     []string{"remove", "pre-commit"},
			Expected: "",
			Error:    false,
		},
	}

	testutil.TestCmd(t, testutil.SetupFunc(func() error {
		outBuf := new(bytes.Buffer)
		errBuf := new(bytes.Buffer)

		rootCmd.SetOut(outBuf)
		rootCmd.SetErr(errBuf)
		rootCmd.SetArgs([]string{"init"})

		if err := rootCmd.Execute(); err != nil {
			return err
		}

		rootCmd.SetArgs([]string{"add", "pre-commit"})

		return rootCmd.Execute()
	}), rootCmd, tests)
}
