/*
Package git provides utilities related to Git hooks and supported scripting
languages for those hooks.

This package includes functionality for retrieving information about Git hooks
and determining the supported languages for writing hooks.

This package is designed to facilitate the use and understanding of Git hooks
in various programming environments.
*/
package git

var (
	// hooks is a map of all available Git hook names.
	//
	// INFO: This map is up to date as of 11/17/2025.
	// SOURCE: https://git-scm.com/docs/githooks
	// NOTE: update the date whenever modifying the map
	hooks = map[string]bool{
		"applypatch-msg":        true,
		"pre-applypatch":        true,
		"post-applypatch":       true,
		"pre-commit":            true,
		"pre-merge-commit":      true,
		"prepare-commit-msg":    true,
		"commit-msg":            true,
		"post-commit":           true,
		"pre-rebase":            true,
		"post-checkout":         true,
		"post-merge":            true,
		"pre-push":              true,
		"pre-receive":           true,
		"update":                true,
		"proc-receive":          true,
		"post-receive":          true,
		"post-update":           true,
		"reference-transaction": true,
		"push-to-checkout":      true,
		"pre-auto-gc":           true,
		"post-rewrite":          true,
		"sendemail-validate":    true,
		"fsmonitor-watchman":    true,
		"p4-changelist":         true,
		"p4-prepare-changelist": true,
		"p4-post-changelist":    true,
		"p4-pre-submit":         true,
		"post-index-change":     true,
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

// CheckHook reports if a specified hook is in supported Git hooks.
func CheckHook(key string) bool {
	_, exist := hooks[key]

	return exist
}

// Hooks returns the complete map of all defined Git hooks.
func Hooks() map[string]bool {
	return hooks
}

// CheckLangSupported reports if a specified language is supported for Git hooks.
func CheckLangSupported(key string) bool {
	_, exist := supportedLangs[key]

	return exist
}

// SupportedLangs returns the map of supported programming languages for Git hooks.
func SupportedLangs() map[string]bool {
	return supportedLangs
}
