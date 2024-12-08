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
	// Makes sure .hkup directory exists in current working directory
	if !util.DoesDirectoryExist(util.HkupDirName) {
		return fmt.Errorf("%s directory does not exist", util.HkupDirName)
	}

	hook := args[0]
	filePath := util.GetHookFilePath(hook)

	// Does not add if hook already exists in .hkup directory
	if util.DoesFileExist(filePath) {
		return fmt.Errorf("%s hook already exists", hook)
	}

	var fileContent string

	// Uses the specified language from lang flag; else default to sh
	if LangFlg != "" {
		// make sure lang is supported
		if isValid := git.CheckLangSupported(LangFlg); !isValid {
			return fmt.Errorf("language not supported: %s", LangFlg)
		}
		fileContent = fmt.Sprintf("#!/usr/bin/env %s\n\n\n\n\n", LangFlg)
	} else {
		fileContent = "#!/bin/sh\n\n\n\n\n"
	}

	file, err := util.CreateFile(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fileContent)
	if err != nil {
		return err
	}

	// Either changes the create file's permissions successful or returns error
	return util.MakeExecutable(filePath)
}
