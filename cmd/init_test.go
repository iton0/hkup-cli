package cmd

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestInitCmd(t *testing.T) {
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
			args: []string{"init"},
			want: "Usage:\n  hkup init [flags]\n\nFlags:\n      --gitdir string     specified path to git directory\n  -h, --help              help for init\n      --worktree string   specified path to working tree\n\n",
			err:  fmt.Errorf("hooksPath already set to .hkup\n"),
		},
		// Add more test cases here if necessary, e.g., for error conditions
	}

	for _, tt := range tests {
		buf.Reset() // Reset the buffer before each command execution
		rootCmd.SetArgs(tt.args)

		err := rootCmd.Execute()

		// Check for expected error
		if (err != nil) != (tt.err != nil) || (err != nil && err.Error() != tt.err.Error()) {
			t.Fatalf("Command failed for args %v: got error %v, want %v", tt.args, err, tt.err)
		}

		got := buf.String()
		if got != tt.want {
			t.Errorf("got last line %q, want %q for args %v", got, tt.want, tt.args)
		}
	}
}
