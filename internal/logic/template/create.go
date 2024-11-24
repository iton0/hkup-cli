package template

import (
	"fmt"
	"path/filepath"

	"github.com/iton0/hkup-cli/internal/git"
	"github.com/iton0/hkup-cli/internal/util"
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

	// template holds the information to create the new template.
	// Info includes:
	//   - git hook name (hook)
	//   - language (lang)
	//   - custom template name (name)
	//   - if to use git hook in the current working directory (useCwd)
	//   - if to copy created template in the current working directory (copyHook)
	//   - if to edit the created template by opening editor (edit)
	template = struct {
		hook, lang, name       string
		useCwd, copyHook, edit bool
	}{}
)

// Create creates a git hook template from a specific git hook.
//
// Returns error if:
//   - issue with creating HkUp config directory or template directory
//   - issue with displaying prompt
//   - issue with creating the template
func Create(cmd *cobra.Command, args []string) error {
	configPath := util.GetConfigDirPath()
	templatePath := util.GetTemplateDirPath()

	// Make the HkUp config directory if it does not exist
	if !util.DoesDirectoryExist(configPath) {
		cmd.Printf("Making HkUp config directory at %s...\n", configPath)

		err := util.CreateDirectory(configPath)
		if err != nil {
			return err
		}

		// Also make the template subdirectory
		err = util.CreateDirectory(templatePath)
		if err != nil {
			return err
		}
	}

	if len(args) == 1 {
		if err := displayPrompt(templatePath, args[0]); err != nil {
			return err
		}
	} else if err := displayPrompt(templatePath); err != nil { // no args given
		return err
	}

	// Either creating the template is successful and returns nil or unsuccessful
	// and returns error
	return createTemplate(templatePath)
}

// createTemplate creates the template based on the given args and flags.
// Returns error if any operation fails.
func createTemplate(templatePath string) error {
	fmt.Println() // Makes the output more distinct in regards to spacing

	// Full path to created template
	createdTemplate := filepath.Join(templatePath, template.name+"#"+template.hook)

	// Copies git hook from current working directory to template directory
	if template.useCwd {
		srcPath := util.GetHookFilePath(template.hook)
		return util.CopyFile(srcPath, createdTemplate) // returns either nil or error
	}

	file, err := util.CreateFile(createdTemplate)
	if err != nil {
		return err
	}
	defer file.Close()

	var fileContent string
	if template.lang == "" { // Default to sh for the template language
		fileContent = "#!/bin/sh\n\n\n\n\n"
	} else {
		fileContent = fmt.Sprintf("#!/usr/bin/env %s\n\n\n\n\n", template.lang)
	}

	_, err = file.WriteString(fileContent)
	if err != nil {
		return err
	}

	if template.edit {
		if err := editTemplate(createdTemplate); err != nil {
			return err
		}
		fmt.Println("Template successfully edited!")
	}

	if template.copyHook {
		dstPath := filepath.Join(util.HkupDirName, template.hook)

		err := util.CopyFile(createdTemplate, dstPath)
		if err != nil {
			return err
		}

		err = util.MakeExecutable(dstPath)
		if err != nil {
			return err
		}

		fmt.Println("Template copied to current working directory.")
	}

	return nil
}

// displayPrompt outputs appropiate prompts based on args and flags of command.
// Returns error if issue with displaying any of the sub prompts.
func displayPrompt(templatePath string, arg ...string) error {
	fmt.Println() // Makes the output more distinct in regards to spacing

	// Takes user provided arg as hook name or asks for it
	if len(arg) == 1 {
		template.hook = arg[0]
		fmt.Printf("Creating template with %s hook...\n\n", template.hook)
	} else if err := displayHookPrompt(); err != nil {
		return err
	}

	// Takes name if name flag used or asks for it
	if TemplateNameFlg != "" {
		if out, err := doesTemplateExist(templatePath, TemplateNameFlg); err != nil {
			return err
		} else if out != "" {
			return fmt.Errorf("template %s already exists", out)
		}
		template.name = TemplateNameFlg
	} else if err := displayNamePrompt(templatePath); err != nil {
		return err
	}

	// Uses the current working directory's git hook if flag given or asks to do so
	// NOTE: If cwd flag is used then utilizes language of that existing hook
	if TemplateCwdFlg {
		if !util.DoesFileExist(filepath.Join(util.HkupDirName, template.hook)) {
			return fmt.Errorf("git hook %s does not exist in the current working directory", template.hook)
		}
		template.useCwd = true
	} else {
		err := displayCwdPrompt()
		if err != nil {
			return err
		}

		// Takes language if lang flag used or asks for it
		if TemplateLangFlg != "" {
			if _, err = git.GetLang(TemplateLangFlg); err != nil {
				return err
			}
			template.lang = TemplateLangFlg
		} else if err = displayLangPrompt(); err != nil {
			return err
		}

		// Copies created template to cwd if copy flag used or asks to do so
		if TemplateCopyFlg {
			template.copyHook = true
		} else if !template.useCwd { // Does not copy what is already in cwd
			if err = displayCopyPrompt(); err != nil {
				return err
			}
		}

		// Created template will be opened in editor or asks to do so
		if TemplateEditFlg {
			template.edit = true
		} else if err = displayEditPrompt(); err != nil {
			return err
		}
	}

	return nil
}

// displayHookPrompt asks for valid git hook name to use for template.
// Returns error is issue with reading response.
func displayHookPrompt() error {
	in, err := util.UserInputPrompt("Git hook name:")
	if err != nil {
		return err
	}

	// Recursively calls this function until supplied with supported git hook
	if _, err = git.GetHook(in); err != nil {
		fmt.Println("Not a supported Git hook. Please try again")
		return displayHookPrompt()
	}

	template.hook = in
	return nil
}

// displayCwdPrompt asks whether to use current working directory's git hook as
// template.
// Returns error is issue with reading response.
func displayCwdPrompt() error {
	// Does not display if the git hook type does not exist in the cwd
	if !util.DoesFileExist(filepath.Join(util.HkupDirName, template.hook)) {
		return nil
	}

	isYes, err := util.YesNoPrompt("Use from current working directory?")
	if err != nil {
		return err
	}

	// useCwd field is false by default so only need to check if "yes"
	template.useCwd = isYes
	return nil
}

// displayLangPrompt asks what language to use for template.
// Returns error if is issue with reading reponse.
func displayLangPrompt() error {
	// Does not display if we are using the existing git hook in cwd
	if template.useCwd {
		return nil
	}

	switch in, err := util.UserInputPrompt("Language (default sh):"); {
	case err != nil:
		return err
	case in == "": // using the default sh as the language for the hook
		return nil
	default:
		// Recursively calls this function until supplied with supported language
		if _, err = git.GetLang(in); err != nil {
			fmt.Println("Not a supported language. Please try again")
			return displayLangPrompt()
		}

		template.lang = in
		return nil
	}
}

// displayNamePrompt asks for the name of the template.
// Returns error if:
//   - issue with reading response
//   - issue with checking config template directory
func displayNamePrompt(templatePath string) error {
	in, err := util.UserInputPrompt("Template Name:")
	if err != nil {
		return err
	}

	if out, err := doesTemplateExist(templatePath, in); err != nil {
		return err
	} else if out != "" { // Keeps asking until given a unique template name
		fmt.Println("Template name already exists. Please try again")
		return displayNamePrompt(templatePath)
	}

	template.name = in
	return nil
}

// displayCopyPrompt asks whether to copy the template to the current working
// directory.
// Returns an error if issue with reading reponse.
func displayCopyPrompt() error {
	isYes, err := util.YesNoPrompt("Copy to current working directory?")
	if err != nil {
		return err
	}

	// copyHook field is false by default so only need to check if "yes"
	template.copyHook = isYes
	return nil
}

// displayEditPrompt asks whether to edit the created template.
// Returns an error if issue with reading reponse.
func displayEditPrompt() error {
	isYes, err := util.YesNoPrompt("Edit template?")
	if err != nil {
		return err
	}

	// edit field is false by default so only need to check if "yes"
	template.edit = isYes
	return nil
}
