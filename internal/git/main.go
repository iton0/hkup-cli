/*
Package git provides utilities related to Git hooks and supported scripting
languages for those hooks.

This package includes functionality for retrieving information about Git hooks
and determining the supported languages for writing hooks.

This package is designed to facilitate the use and understanding of Git hooks
in various programming environments.
*/
package git

import (
	"fmt"
)

var (
	// hooks is a map of Git hook names to their respective section of the
	// Git hooks documentation site.
	//
	// INFO: This map is up to date as of 11/19/2024.
	// Source: https://git-scm.com/docs/githooks
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

	// supportedLangs is a map indicating which programming languages are supported
	// for Git hooks.
	supportedLangs = map[string]bool{
		"sh":     true,
		"bash":   true,
		"python": true,
		"ruby":   true,
		"node":   true,
		"perl":   true,
		"php":    true,
	}
)

// GetHookUrl retrieves URL section of git doc site for specified Git hook.
// Returns URL section of specified hook or empty string if hook not supported.
func GetHookUrl(key string) string {
	val, exist := hooks[key]
	if !exist {
		return ""
	}

	return val
}

// Hooks returns the complete map of all defined Git hooks.
func Hooks() map[string]string {
	return hooks
}

// CheckLangSupported reports if a specified language is supported for Git hooks.
// Returns boolean and error if the language is not recognized.
func CheckLangSupported(key string) (bool, error) {
	if _, exist := supportedLangs[key]; !exist {
		return false, fmt.Errorf("language not supported: %s", key)
	}

	return true, nil
}

// SupportedLangs returns the map of supported programming languages for Git hooks.
func SupportedLangs() map[string]bool {
	return supportedLangs
}
