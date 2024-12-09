package cmd

import (
	"bytes"
	"fmt"
	"testing"
)

// TODO: update this test function appropiately
// TestTemplateCmd tests use cases for the hkup template command.
func TestTemplateCmd(t *testing.T) {
	var buf bytes.Buffer
	rootCmd.SetOut(&buf)

	tests := []struct {
		args []string
		want string
		err  error
	}{
		{
			args: []string{"template", "create", "test"},
			want: "",
			err:  fmt.Errorf("invalid argument \"test\" for \"hkup template create\""),
		},
		// Add more test cases here if necessary, e.g., for error conditions
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
