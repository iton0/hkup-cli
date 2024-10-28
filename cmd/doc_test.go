package cmd

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestDocCmd(t *testing.T) {
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
			args: []string{"doc", "test"},
			want: "Usage:\n  hkup doc <hook-name> [flags]\n\nAliases:\n  doc, docs\n\nFlags:\n  -h, --help   help for doc\n\n",
			err:  fmt.Errorf("invalid argument \"test\" for \"hkup doc\""),
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
			t.Errorf("got %q, want %q for args %v", got, tt.want, tt.args)
		}
	}
}
