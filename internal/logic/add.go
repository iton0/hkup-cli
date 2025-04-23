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
// programming language in the designated .hkup directory. Returns error if any
// of the steps fail above.
func Add(_ *cobra.Command, args []string) error {
	_, err := isBareRepo(".")
	if err != nil { // Current working directory is not a git repository at all
		return err
	}

	hook := args[0]
	filePath := util.GetHookFilePath(hook)

	if util.DoesFileExist(filePath) {
		return fmt.Errorf("%s hook already exists", hook)
	}

	var sheBangLine string

	if LangFlg != "" {
		if isValid := git.CheckLangSupported(LangFlg); !isValid {
			return fmt.Errorf("language not supported: %s", LangFlg)
		}
		sheBangLine = fmt.Sprintf("#!/usr/bin/env %s", LangFlg)
	} else {
		sheBangLine = "#!/bin/sh"
	}

	file, err := util.CreateFile(filePath)
	defer file.Close()
	if err != nil {
		return err
	}

	_, err = file.WriteString(sheBangLine)
	if err != nil {
		return err
	}

	return util.MakeExecutable(filePath)
}
