/*
Package util provides utility functions for file and directory operations, including creating folders and files, checking existence of files and directories, and converting map keys to slices.

This package is designed to simplify common file and directory management tasks for hkup related commands in the [internal/logic] package.
*/
package util

import (
	"fmt"
	"os"
)

// CreateFolder creates a new directory at the specified path. Returns an error if the operation fails.
func CreateFolder(dirPath string) error {
	if err := os.Mkdir(dirPath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dirPath, err)
	}
	return nil
}

// CreateFile creates a new file in the specified directory with the given name. Returns an error if the operation fails.
func CreateFile(dirPath string, name string) error {
	filePath := dirPath + name
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filePath, err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	return nil
}

// DoesDirectoryExist reports if a directory exists at the specified path.
func DoesDirectoryExist(dirPath string) bool {
	info, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// DoesFileExist reports if a file exists at the specified path.
func DoesFileExist(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// ConvertMapKeysToSlice converts the keys of a map into a returned slice of strings.
func ConvertMapKeysToSlice[T comparable](m map[string]T) []string {
	var keys []string

	for key := range m {
		keys = append(keys, key)
	}

	return keys
}
