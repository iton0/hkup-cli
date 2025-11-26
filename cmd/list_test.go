package cmd

// import (
// 	"testing"
//
// 	"github.com/iton0/hkup-cli/v2/internal/testutil"
// )
//
// func TestListCmd(t *testing.T) {
// 	tests := []testutil.Test{
// 		{
// 			Name: "Valid Arg for Hooks",
// 			Args: []string{"list", "hook"},
// 			// FIXME: list order is not deterministic need to figure out how to
// 			// make it such or look at other values validate
// 			Expected: "post-applypatch\n pre-rebase\n update\n post-update\n fsmonitor-watchman\n pre-applypatch\n prepare-commit-msg\n commit-msg\n post-merge\n reference-transaction\n pre-auto-gc\n post-rewrite\n sendemail-validate\n post-checkout\n post-receive\n push-to-checkout\n p4-changelist\n p4-prepare-changelist\n post-index-change\n pre-commit\n pre-merge-commit\n post-commit\n pre-push\n pre-receive\n proc-receive\n p4-post-changelist\n p4-pre-submit\n applypatch-msg",
// 			Error:    false,
// 		},
// 		{
// 			Name: "Valid Arg for Languages",
// 			Args: []string{"list", "lang"},
// 			// FIXME: list order is not deterministic need to figure out how to
// 			// make it such or look at other values validate
// 			Expected: "node\n perl\n php\n sh\n bash\n python\n ruby",
// 			Error:    false,
// 		},
// 	}
//
// 	testutil.TestCmd(t, testutil.SetupFunc(func() error {
// 		return nil
// 	}), rootCmd, tests)
// }
