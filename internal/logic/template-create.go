package logic

import (
	"fmt"
	"path/filepath"

	"github.com/iton0/hkup-cli/v2/internal/git"
	"github.com/iton0/hkup-cli/v2/internal/util"
	"github.com/spf13/cobra"
)

var (
	// TemplateLangFlg is an optional flag indicating the language to use.
	TemplateLangFlg string

	// TemplateNameFlg is an optional flag that holds the template name to be
	// prepended to the template name.
	TemplateNameFlg string

	// TemplateCwdFlg is an optional flag indicating to use the git hook from
	// the current working directory.
	TemplateCwdFlg bool

	// TemplateCopyFlg is an optional flag indicating to copy the created
	// template to the current working directory.
	TemplateCopyFlg bool

	// TemplateEditFlg is an optional flag indicating to edit the template.
	TemplateEditFlg bool

	// templatePtr is a pointer that holds information to create a new templatePtr.
	// Info includes:
	//   - git hook name (hook)
	//   - language (lang)
	//   - custom templatePtr name (name)
	//   - if to use git hook in the current working directory (useCwd)
	//   - if to copy created templatePtr in the current working directory (copyHook)
	//   - if to edit the created templatePtr by opening editor (edit)
	templatePtr = &struct {
		hook, lang, name       string
		useCwd, copyHook, edit bool
	}{}
)

// TemplateCreate creates a git hook template from a specific git hook.
//
// Returns error if:
//   - issue with creating HkUp config directory or template directory
//   - issue with displaying prompt
//   - issue with creating the template
func TemplateCreate(cmd *cobra.Command, args []string) error {
	configPath := util.GetConfigDirPath()
	templatePath := util.GetTemplateDirPath()

	switch {
	case !util.DoesDirectoryExist(configPath):
		cmd.Printf("Making HkUp config directory at %s...\n", configPath)

		// Creates both the Hkup config directory and its template subdirectory
		if err := util.CreateDirectory(templatePath); err != nil {
			return err
		}

		file, err := util.CreateFile(util.GetConfigFilePath())
		if err != nil {
			return err
		}
		defer file.Close()
	case !util.DoesDirectoryExist(templatePath):
		cmd.Printf("Making HkUp template directory at %s...\n", templatePath)

		if err := util.CreateDirectory(templatePath); err != nil {
			return err
		}
	}

	if len(args) == 1 {
		if err := displayPrompt(templatePath, args[0]); err != nil {
			return err
		}
	} else if err := displayPrompt(templatePath); err != nil {
		return err
	}

	return createTemplate(templatePath)
}

// createTemplate creates the template based on the given args and flags.
// Returns error if any operation fails.
func createTemplate(templatePath string) error {
	fmt.Println() // Makes the output more distinct in regards to spacing

	createdTemplateFullPath := filepath.Join(templatePath, templatePtr.name+"#"+templatePtr.hook)

	if templatePtr.useCwd {
		srcPath := util.GetHookFilePath(templatePtr.hook)
		return util.CopyFile(srcPath, createdTemplateFullPath)
	}

	file, err := util.CreateFile(createdTemplateFullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	var fileContent string
	if templatePtr.lang == "" {
		fileContent = "#!/bin/sh"
	} else {
		fileContent = fmt.Sprintf("#!/usr/bin/env %s", templatePtr.lang)
	}

	_, err = file.WriteString(fileContent)
	if err != nil {
		return err
	}

	if templatePtr.edit {
		if err := editTemplate(createdTemplateFullPath); err != nil {
			return err
		}
		fmt.Println("Template successfully edited!")
	}

	if templatePtr.copyHook {
		dstPath := filepath.Join(util.HkupDirName, templatePtr.hook)

		err := util.CopyFile(createdTemplateFullPath, dstPath)
		if err != nil {
			return err
		}

		err = util.MakeExecutable(dstPath)
		if err != nil {
			return err
		}

		fmt.Println("Template copied to current working directory.")
	}

	fmt.Printf("Template created at %s\n", createdTemplateFullPath)
	return nil
}

// displayPrompt outputs appropriate prompts based on args and flags of command.
// Returns error if issue with displaying any of the sub prompts.
func displayPrompt(templatePath string, arg ...string) error {
	fmt.Println() // Makes the output more distinct in regards to spacing

	if len(arg) == 1 {
		templatePtr.hook = arg[0]
		fmt.Printf("Creating template with %s hook...\n\n", templatePtr.hook)
	} else if err := displayHookPrompt(); err != nil {
		return err
	}

	if TemplateNameFlg != "" {
		if out, err := doesTemplateExist(templatePath, TemplateNameFlg); err != nil {
			return err
		} else if out != "" {
			return fmt.Errorf("template %s already exists", out)
		}
		templatePtr.name = TemplateNameFlg
	} else if err := displayNamePrompt(templatePath); err != nil {
		return err
	}

	if TemplateCwdFlg {
		if !util.DoesFileExist(filepath.Join(util.HkupDirName, templatePtr.hook)) {
			return fmt.Errorf("git hook %s does not exist in the current working directory", templatePtr.hook)
		}
		templatePtr.useCwd = true
	} else {
		err := displayCwdPrompt()
		if err != nil {
			return err
		}

		if TemplateLangFlg != "" {
			if isValid := git.CheckLangSupported(TemplateLangFlg); !isValid {
				return fmt.Errorf("language not supported: %s", TemplateLangFlg)
			}
			templatePtr.lang = TemplateLangFlg
		} else if err = displayLangPrompt(); err != nil {
			return err
		}

		if TemplateCopyFlg {
			templatePtr.copyHook = true
		} else if !templatePtr.useCwd {
			if err = displayCopyPrompt(); err != nil {
				return err
			}
		}

		if TemplateEditFlg {
			templatePtr.edit = true
		} else if err = displayEditPrompt(); err != nil {
			return err
		}
	}

	return nil
}

// displayHookPrompt asks for valid git hook name to use for template.
// Returns error if issue with reading response or after 3 incorrect attempts.
func displayHookPrompt(attempts ...int) error {
	if len(attempts) == 0 {
		attempts = append(attempts, 0)
	}

	attempt := attempts[0]
	if attempt == 3 {
		return fmt.Errorf("3 incorrect attempts")
	}

	in, err := util.UserInputPrompt("Git hook name:")
	if err != nil {
		return err
	}

	if isValid := git.CheckHook(in); !isValid {
		attempt++
		fmt.Println("Not a supported Git hook. Please try again")
		return displayHookPrompt(attempt)
	}

	templatePtr.hook = in
	return nil
}

// displayCwdPrompt asks whether to use current working directory's git hook as
// template.
// Returns error if issue with reading response or after 3 incorrect attempts.
func displayCwdPrompt() error {
	if !util.DoesFileExist(filepath.Join(util.HkupDirName, templatePtr.hook)) {
		return nil
	}

	isYes, err := util.YesNoPrompt("Use from current working directory?")
	if err != nil {
		return err
	}

	templatePtr.useCwd = isYes
	return nil
}

// displayLangPrompt asks what language to use for template.
// Returns error if issue with reading response or after 3 incorrect attempts.
func displayLangPrompt(attempts ...int) error {
	if len(attempts) == 0 {
		attempts = append(attempts, 0)
	}

	attempt := attempts[0]
	if attempt == 3 {
		return fmt.Errorf("3 incorrect attempts")
	}

	if templatePtr.useCwd {
		return nil
	}

	if in, err := util.UserInputPrompt("Language (default sh):"); err != nil {
		return err
	} else if in == "" {
		return nil
	} else {
		if isValid := git.CheckLangSupported(in); !isValid {
			attempt++
			fmt.Println("Not a supported language. Please try again")
			return displayLangPrompt(attempt)
		}

		templatePtr.lang = in
		return nil
	}
}

// displayNamePrompt asks for the name of the template.
// Returns error if:
//   - issue with reading response
//   - issue with checking config template directory
//   - 3 incorrect name attempts
func displayNamePrompt(templatePath string, attempts ...int) error {
	if len(attempts) == 0 {
		attempts = append(attempts, 0)
	}

	attempt := attempts[0]
	if attempt == 3 {
		return fmt.Errorf("3 incorrect attempts")
	}

	in, err := util.UserInputPrompt("Template Name:")
	if err != nil {
		return err
	}

	if out, err := doesTemplateExist(templatePath, in); err != nil {
		return err
	} else if out != "" {
		attempt++
		fmt.Println("Template name already exists. Please try again")
		return displayNamePrompt(templatePath, attempt)
	}

	templatePtr.name = in
	return nil
}

// displayCopyPrompt asks whether to copy the template to the current working
// directory.
// Returns an error if issue with reading response.
func displayCopyPrompt() error {
	isYes, err := util.YesNoPrompt("Copy to current working directory?")
	if err != nil {
		return err
	}

	templatePtr.copyHook = isYes
	return nil
}

// displayEditPrompt asks whether to edit the created template.
// Returns an error if issue with reading response.
func displayEditPrompt() error {
	isYes, err := util.YesNoPrompt("Edit template?")
	if err != nil {
		return err
	}

	templatePtr.edit = isYes
	return nil
}
