// Package testutil provides utilities for testing Cobra CLI commands.
package testutil

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

// Test represents a single test case for a Cobra command.
type Test struct {
	Name     string
	Args     []string
	Expected string
	Error    bool
}

// SetupFunc is a function signature for command-specific setup logic
// that is executed before running the test cases.
type SetupFunc func() error

// TestCmd executes a series of tests against a given Cobra command.
//
// It sets up a temporary directory with an initialized Git repository,
// runs an optional setup function, and then iterates through the provided
// test cases, executing the command and verifying the output or error.
func TestCmd(t *testing.T, setup SetupFunc, cmd *cobra.Command, tests []Test) {
	t.Helper()

	currDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("could not get current working directory: %v", err)
	}

	dname, err := os.MkdirTemp("", "clitest-")
	if err != nil {
		t.Fatalf("could not create temporary directory: %v", err)
	}

	t.Cleanup(func() {
		if cerr := os.Chdir(currDir); cerr != nil {
			t.Errorf("failed to change back to original directory %s: %v", currDir, cerr)
		}

		if rerr := os.RemoveAll(dname); rerr != nil {
			t.Errorf("failed to remove temporary directory %s: %v", dname, rerr)
		}
	})

	if err := os.Chdir(dname); err != nil {
		t.Fatalf("failed to change to temporary directory: %v", err)
	}

	if err := initializeGitRepo(dname); err != nil {
		t.Fatalf("Error initializing Git repository: %v", err)
		return
	}

	if err := setup(); err != nil {
		t.Fatalf("failed to run command specific setup: %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			outBuf := new(bytes.Buffer)
			errBuf := new(bytes.Buffer)

			cmd.SetOut(outBuf)
			cmd.SetErr(errBuf)
			cmd.SetArgs(tt.Args)

			err := cmd.Execute()

			if tt.Error {
				if err == nil {
					t.Fatal("expected an error, but command executed successfully")
				}

				errorOutput := strings.TrimSpace(errBuf.String())

				if !strings.Contains(errorOutput, tt.Expected) {
					t.Errorf("expected error output to contain %q, but got %q", tt.Expected, errBuf.String())
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error occurred: %v\nStderr: %s", err, errBuf.String())
				}

				output := strings.TrimSpace(outBuf.String())
				expectedTrimmed := strings.TrimSpace(tt.Expected)

				if output != expectedTrimmed {
					t.Errorf("output mismatch:\nExpected: %q\nGot: %q", expectedTrimmed, output)
				}
			}
		})
	}
}

// initializeGitRepo initializes an empty Git repository in the given directory.
func initializeGitRepo(dir string) error {
	cmd := exec.Command("git", "init")

	cmd.Dir = dir

	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run git init: %w", err)
	}

	return nil
}
