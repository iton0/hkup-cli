package logic

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/iton0/hkup-cli/internal/git"
	"github.com/iton0/hkup-cli/internal/util"
	"github.com/spf13/cobra"
)

var (
	// Lang is an optional flag indicating the programming language to use for the hook script. Defaults to sh.
	Lang string
)

// Add adds a new Git hook with the specified name and optional programming language.
// It creates a new file in the designated `.hkup` directory, setting the appropriate shebang line based on the provided language.
// Returns an error if any of the steps fail, including directory existence, file creation, or permission setting.
func Add(cmd *cobra.Command, args []string) error {
	var sheBangLine = "#!/bin/sh\n\n\n\n\n"
	hook := args[0]

	if Lang != "" {
		if _, err := git.GetLang(Lang); err != nil {
			return err
		}
		sheBangLine = fmt.Sprintf("#!/usr/bin/env %s\n\n\n\n\n", Lang)
	}

	if !util.DoesDirectoryExist(FullPath) {
		return fmt.Errorf("failed running \"hkup add\"\n%s does not exist", FullPath)
	}

	filePath := filepath.Join(FullPath, hook)

	if util.DoesFileExist(filePath) {
		return fmt.Errorf("%s hook already exists", hook)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	_, err = file.WriteString(sheBangLine)
	if err != nil {
		return fmt.Errorf("failed writing to file: %w", err)
	}

	err = os.Chmod(filePath, 0755)
	if err != nil {
		return fmt.Errorf("failed changing permissions of file: %w", err)
	}

	return nil
}
