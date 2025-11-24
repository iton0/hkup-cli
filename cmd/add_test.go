package cmd

import (
	"bytes"
	"testing"

	"github.com/iton0/hkup-cli/v2/internal/testutil"
)

func TestAddCmd(t *testing.T) {
	tests := []testutil.Test{
		{
			Name:     "Valid Arg",
			Args:     []string{"add", "pre-commit"},
			Expected: "",
			Error:    false,
		},
		{
			Name:     "Valid Arg w/ Language flag",
			Args:     []string{"add", "--lang", "python", "post-commit"},
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

		return rootCmd.Execute()
	}), rootCmd, tests)
}
