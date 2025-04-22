/*
Package logic provides functionality for managing Git hooks, including commands
to add, remove, and list hooks.
This package utilizes the cobra library for command-line interaction and is
implemented in the respective commands of the [github.com/iton0/hkup-cli/cmd] package.

Commands:
  - Init: Initializes HkUp.
  - End: Resets the local hooksPath variable.
  - Add: Adds a new Git hook with the specified name and optional programming language.
  - Remove: Removes an existing Git hook with the specified name.
  - List: Lists all available Git hooks or supported Git hook languages.
  - Doc: Opens browser for specified Git hook documentation.
  - Root: Wraps git-related clone commands for easier initialization of HkUp.
  - Status: Shows if HkUp is active for current working directory.
*/
package logic

// NOTE: This file is for documentation purposes and should be kept empty.
