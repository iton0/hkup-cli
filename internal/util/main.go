/*
Package util provides utility functions for tasks and operations such as:
  - terminal prompt creation
  - file/directory operations
  - retrieval of related system information
  - getting/setting of HkUp configuration settings
  - mutation of related data structures

Additionally, this package holds all constant values used throughout the
application such as:
  - HkUp directory name

This package is designed to abstract the above values, tasks, and operations to
be reusable throughout the HkUp application.
*/
package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	// HkupDirName defines HkUp directory name within current working directory.
	HkupDirName = ".hkup"
)

// CreateDirectory makes a new directory at the specified path.
// Returns an error if the operation fails.
func CreateDirectory(path string) error {
	return os.Mkdir(path, os.ModePerm)
}

// CreateFile makes a new file in the specified file path name.
// Returns pointer to the new file and an error if the operation fails.
// NOTE: CreateFile does not close the file.
func CreateFile(path string) (*os.File, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// DoesDirectoryExist reports if a directory exists at the specified path.
func DoesDirectoryExist(path string) bool {
	if info, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return info.IsDir()
	}
}

// DoesFileExist reports if a file exists at the specified path.
func DoesFileExist(path string) bool {
	if info, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return !info.IsDir()
	}
}

// GetConfigDirPath returns the available HkUp config directory path.
func GetConfigDirPath() (configPath string) {
	if xdgVar, exist := os.LookupEnv("XDG_CONFIG_HOME"); exist && xdgVar != "" {
		configPath = filepath.Join(xdgVar, "hkup")
	} else {
		configPath = filepath.Join(os.Getenv("HOME"), ".config", "hkup")
	}

	return configPath
}

// GetConfigFilePath returns the HkUp file path that holds configuration settings.
func GetConfigFilePath() string {
	return filepath.Join(GetConfigDirPath(), ".hkupconfig")
}

// GetTemplateDirPath returns the HkUp config template directory path.
func GetTemplateDirPath() string {
	return filepath.Join(GetConfigDirPath(), "templates")
}

// CopyFile copies a file (without overwriting) from src file path to dest file path.
// Returns error if:
//   - destination path exists
//   - issue with any steps of copying
func CopyFile(src, dst string) error {
	if _, err := os.Stat(dst); err == nil {
		return fmt.Errorf("destination file already exists: %s", dst)
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("failed to check if destination file exists: %w", err)
	}

	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	if _, err = io.Copy(dstFile, srcFile); err != nil {
		return err
	}

	return nil
}

// MakeExecutable makes the filePath executable.
// Returns error if issue with making executable.
func MakeExecutable(filePath string) error {
	return os.Chmod(filePath, 0755)
}

// ConvertMapKeysToSlice transforms the map string keys into a returned slice
// of strings.
func ConvertMapKeysToSlice[T comparable](m map[string]T) []string {
	keys := []string{}

	for key := range m {
		keys = append(keys, key)
	}

	return keys
}

// YesNoPrompt displays the specified prompt message to the user and asks for a
// yes/no response.
// Returns boolean and error if issue occurred during the input process.
func YesNoPrompt(prompt string) (bool, error) {
	fmt.Print(prompt + "(Y/n): ")

	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return false, fmt.Errorf("failed to read response")
	}

	response := strings.TrimSpace(scanner.Text())

	// Pressing Enter key is equivalent to yes
	if response == "" || response == "y" || response == "Y" {
		return true, nil
	}

	return false, nil
}

// UserInputPrompt prompts the user with the specified message and waits for
// the user to enter a response.
// Returns reponse and error if issue occurred during the input process.
func UserInputPrompt(prompt string) (string, error) {
	fmt.Print(prompt + " ")

	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return "", fmt.Errorf("failed to read response")
	}

	return strings.TrimSpace(scanner.Text()), nil
}

// GetEditor makes best effort to find default editor for HkUp.
// Returns editor name if found and error if issue with searching for editor.
func GetEditor() (string, error) {
	// Check the HkUp config file
	editor, err := GetTOMLValue(GetConfigFilePath(), "editor")
	if err != nil {
		return "", err
	} else if editor != "" {
		return editor, nil
	}

	// Check in global gitconfig file
	if out, err := exec.Command("git", "config", "--global", "core.editor").CombinedOutput(); err != nil {
		return "", err
	} else if len(out) != 0 {
		// The out has a newline character at the end so take elements up until the
		// "\" of the "\n"
		return string(out[0:(len(out) - 1)]), nil // Converts byte slice into string
	}

	// Check for EDITOR var
	if editor, exist := os.LookupEnv("EDITOR"); exist && editor != "" {
		return editor, nil
	}

	return "", fmt.Errorf("failed to find an editor")
}

// GetTOMLValue gets the value of a specific key from a flat TOML file.
// Returns value and error if issue with opening or reading file.
func GetTOMLValue(filePath, key string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", nil
	}
	defer file.Close()

	// Scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		// Skip comments or empty lines
		if len(line) == 0 || line[0] == '#' || line[0] == ';' {
			continue
		}

		// Split the line into key and value
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue // Skip malformed lines
		}

		keyInFile := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Remove quotes around string values (if any)
		if len(value) > 1 && value[0] == '"' && value[len(value)-1] == '"' {
			value = value[1 : len(value)-1]
		}

		// If the current line matches the key you're looking for, return the value
		if keyInFile == key {
			return value, nil
		}
	}

	// Handle the case where the key was not found
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", fmt.Errorf("%s is not a valid key", key) // Returns empty string if key not found
}

// SetTOMLValue modifies the value of a key in a flat TOML file.
// Returns error if key not found or issue with reading or wriiting to file.
func SetTOMLValue(filePath, key, newValue string) error {
	// Open the TOML file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	updatedLines := []string{}
	var keyFound bool // defaults to false

	// Scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines or comments
		if len(line) == 0 || line[0] == '#' || line[0] == ';' {
			updatedLines = append(updatedLines, line)
			continue
		}

		// Split the line into key and value
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			keyInFile := strings.TrimSpace(parts[0])

			// If the key matches, update the value
			if keyInFile == key {
				line = fmt.Sprintf("%s = %s", keyInFile, newValue)
				keyFound = true
			}
		}

		updatedLines = append(updatedLines, line)
	}

	// If the key was not found, return an error
	if !keyFound {
		return fmt.Errorf("key '%s' not found", key)
	}

	// Write the updated content back to the file
	err = os.WriteFile(filePath, []byte(strings.Join(updatedLines, "\n")), 0644)
	if err != nil {
		return err
	}

	return nil
}
