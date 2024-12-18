package logic

import (
	"fmt"

	"github.com/iton0/hkup-cli/internal/git"
	"github.com/iton0/hkup-cli/internal/util"
	"github.com/spf13/cobra"
)

// LangFlg is an optional flag indicating the programming language to use for
// the hook script.
var LangFlg string

// Add creates a new Git hook with the specified git hook name and optional
// programming language in the designated .hkup directory.
//
// Returns error if any of the steps fail above.
func Add(cmd *cobra.Command, args []string) error {
	isBare, err := isBareRepo(".")
	if err != nil { // Current working directory is not a git repository at all
		return err
	}

	// Tries to create .hkup directory if it does not exist and current working
	// directory is not a bare git repository
	if !util.DoesDirectoryExist(util.HkupDirName) && !isBare {
		if err := util.CreateDirectory(util.HkupDirName); err != nil {
			return err
		}

		cmd.Printf("Initialized hkup directory at %s\n", util.HkupDirName)
	}

	hook := args[0]
	filePath := util.GetHookFilePath(hook)

	// Does not add if hook already exists in .hkup directory
	if util.DoesFileExist(filePath) {
		return fmt.Errorf("%s hook already exists", hook)
	}

	var sheBangLine string

	// Uses the specified language from lang flag; else default to sh
	if LangFlg != "" {
		// Makes sure lang is supported
		if isValid := git.CheckLangSupported(LangFlg); !isValid {
			return fmt.Errorf("language not supported: %s", LangFlg)
		}
		sheBangLine = fmt.Sprintf("#!/usr/bin/env %s", LangFlg)
	} else {
		sheBangLine = "#!/bin/sh"
	}

	// Creates the git hook file
	file, err := util.CreateFile(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Writes the language shebang line to the created file from above
	_, err = file.WriteString(sheBangLine)
	if err != nil {
		return err
	}

	// Either changes the create file's permissions successful or returns error
	return util.MakeExecutable(filePath)
}
