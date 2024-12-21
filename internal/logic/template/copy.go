package template

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/iton0/hkup-cli/internal/util"
	"github.com/spf13/cobra"
)

// Copy copies a git hook template to the .hkup directory.
//
// Returns error if:
//   - HkUp config directory does not exist
//   - .hkup directory does not exist in current working directory
//   - arg is not valid template name
//   - issue with copying template to .hkup directory
func Copy(cmd *cobra.Command, args []string) error {
	configPath := util.GetConfigDirPath()
	var templateName string

	switch {
	case !util.DoesDirectoryExist(configPath):
		return fmt.Errorf("%s directory does not exist", configPath)
	case !util.DoesDirectoryExist(util.HkupDirName):
		return fmt.Errorf("%s directory does not exist in current working directory", util.HkupDirName)
	default:
		templateName = args[0]
	}

	templatePath := util.GetTemplateDirPath()
	file, err := doesTemplateExist(templatePath, templateName)
	switch {
	case err != nil:
		return err
	case file == "":
		return fmt.Errorf("not a valid arg \"%s\" for \"hkup template copy\"", templateName)
	default:
		return performCopy(file)
	}
}

// doesTemplateExist checks if any file in the directory specified by templatePath
// starts with the given prefix (template name).
//
// Returns:
//   - The full file path of the first file that matches the given prefix, or an empty string if no match is found.
//   - An error if there is an issue reading the directory.
func doesTemplateExist(templatePath, name string) (string, error) {
	files, err := os.ReadDir(templatePath)
	if err != nil {
		return "", err
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), name) {
			return filepath.Join(templatePath, file.Name()), nil
		}
	}

	return "", nil
}

// performCopy copies the template file to the current working directory with
// appropriate git hook name.
//
// Returns error if:
//   - template does not follow naming convetion
//   - issues with copying or making executable
func performCopy(file string) error {
	cleanPath, err := cleanPath(file)
	if err != nil {
		return err
	}

	dstPath := util.GetHookFilePath(cleanPath)

	err = util.CopyFile(file, dstPath)
	if err != nil {
		return err
	}

	return util.MakeExecutable(dstPath)
}

// cleanPath takes the template file path and returns the substring of the valid
// git hook file name. If the template path does not follow the convention of
// template path naming it will return an empty string and error.
//
// NOTE: This is automatically done by HkUp when using the CLI but user may
// want to manual add a git hook template to the HkUp config template directory.
//
// The convention should follow the custom name of the hook followed by a "#"
// and then the proper git hook name.
//
// Valid Naming Convention: [custom-name]#[hook-name]
//   - ex). foo#post-commit
func cleanPath(path string) (string, error) {
	if idx := strings.LastIndex(path, "#"); idx != -1 {
		return path[idx+1:], nil
	}

	return "", fmt.Errorf("template file name must follow convention of \"<template-name>#<hook-name>\"")
}
