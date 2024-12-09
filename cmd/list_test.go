package cmd

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// TestListCmd tests use cases for the hkup list command.
func TestListCmd(t *testing.T) {
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
			args: []string{"list"},
			want: "",
			err:  fmt.Errorf("accepts 1 arg(s), received 0"),
		},
		{
			args: []string{"list", "test"},
			want: "",
			err:  fmt.Errorf("invalid argument \"test\" for \"hkup list\""),
		},
	}

	for _, tt := range tests {
		buf.Reset() // Reset the buffer before each command execution
		rootCmd.SetArgs(tt.args)

		err := rootCmd.Execute()

		// Check for expected error
		if err.Error() != tt.err.Error() && err != nil {
			t.Fatalf("Command failed for args %v: got error %v, want %v", tt.args, err, tt.err)
		}

		got := buf.String()
		if tt.want != "" && got != tt.want {
			t.Errorf("got output %q, want %q for args %v", got, tt.want, tt.args)
		}
	}
}
