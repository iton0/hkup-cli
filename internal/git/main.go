/*
Package git provides utilities related to Git hooks and supported scripting languages for those hooks.

This package includes functionality for retrieving information about Git hooks and determining the supported languages for writing hooks.

This package is designed to facilitate the use and understanding of Git hooks in various programming environments.
*/
package git

import (
	"fmt"
)

// HookDocSite is a constant of the base URL for the Git hooks documentation.
const HookDocSite = "https://git-scm.com/docs/githooks#"

var (
	// hooks is a map of Git hook names to their corresponding sections in the Git hooks documentation. This map is kept up to date as of 10/24/2024.
	// source: https://git-scm.com/docs/githooks
	hooks = map[string]string{
		"applypatch-msg":        "_applypatch_msg",
		"pre-applypatch":        "_pre_applypatch",
		"post-applypatch":       "_post_applypatch",
		"pre-commit":            "_pre_commit",
		"pre-merge-commit":      "_pre_merge_commit",
		"prepare-commit-msg":    "_prepare_commit_msg",
		"commit-msg":            "_commit_msg",
		"post-commit":           "_post_commit",
		"pre-rebase":            "_pre_rebase",
		"post-checkout":         "_post_checkout",
		"post-merge":            "_post_merge",
		"pre-push":              "_pre_push",
		"pre-receive":           "pre-receive",
		"update":                "update",
		"proc-receive":          "proc-receive",
		"post-receive":          "post-receive",
		"post-update":           "post-update",
		"reference-transaction": "_reference_transaction",
		"push-to-checkout":      "_push_to_checkout",
		"pre-auto-gc":           "_pre_auto_gc",
		"post-rewrite":          "_post_rewrite",
		"sendemail-validate":    "_sendemail_validate",
		"fsmonitor-watchman":    "_fsmonitor_watchman",
		"p4-changelist":         "_p4_changelist",
		"p4-prepare-changelist": "_p4_prepare_changelist",
		"p4-post-changelist":    "_p4_post_changelist",
		"p4-pre-submit":         "_p4_pre_submit",
		"post-index-change":     "_post_index_change",
	}

	// supportedLangs is a map indicating which programming languages are supported for Git hooks, excluding the default bash.
	supportedLangs = map[string]bool{
		"python": true,
		"ruby":   true,
		"node":   true,
		"perl":   true,
		"php":    true,
	}
)

// GetHook retrieves the URL section of the documentation for a specified Git hook. Returns an error if the hook is not found.
func GetHook(key string) (string, error) {
	value, exists := hooks[key]
	if !exists {
		return "", fmt.Errorf("hook not found: %s", key)
	}
	return value, nil
}

// Hooks returns the complete map of all defined Git hooks.
func Hooks() map[string]string {
	return hooks
}

// GetLang reports if a specified language is supported for Git hooks. Returns an error if the language is not recognized.
func GetLang(key string) (bool, error) {
	value, exists := supportedLangs[key]
	if !exists {
		return false, fmt.Errorf("language not supported: %s", key)
	}
	return value, nil
}

// SupportedLangs returns the map of supported programming languages for Git hooks.
func SupportedLangs() map[string]bool {
	return supportedLangs
}
