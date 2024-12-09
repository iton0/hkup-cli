package cmd

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// TestAddCmd tests use cases for the hkup add command.
func TestAddCmd(t *testing.T) {
	var buf bytes.Buffer
	rootCmd.SetOut(&buf)

	// Change directory to the parent
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("could not get current working directory: %v", err)
	}
	defer os.Chdir(originalDir) // Restore original directory after test

	if err := os.Chdir(filepath.Join(originalDir, "..")); err != nil {
		t.Fatalf("could not change to parent directory: %v", err)
	}

	tests := []struct {
		args []string
		want string
		err  error
	}{
		{
			[]string{"add", "test"},
			"",
			fmt.Errorf("invalid argument \"test\" for \"hkup add\""),
		},
		{
			[]string{"add", "pre-commit"},
			"",
			fmt.Errorf("pre-commit hook already exists"),
		},
		{
			[]string{"add", "fsmonitor-watchman", "--lang", "bash"},
			"",
			nil,
		},
		// Add more test cases here if necessary, e.g., for error conditions
	}

	for _, tt := range tests {
		buf.Reset() // Reset the buffer before each command execution
		rootCmd.SetArgs(tt.args)

		err := rootCmd.Execute()

		// Check for expected error
		if err != nil && err.Error() != tt.err.Error() {
			t.Fatalf("Command failed for args %v: got error %v, want %v", tt.args, err, tt.err)
		}

		got := buf.String()
		if tt.want != "" && got != tt.want {
			t.Errorf("got output %q, want %q for args %v", got, tt.want, tt.args)
		}
	}
}
