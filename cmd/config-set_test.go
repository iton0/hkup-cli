package cmd

// import (
// 	"testing"
//
// 	"github.com/iton0/hkup-cli/v2/internal/testutil"
// )
//
// // FIXME: how to test without changing users config file
// func TestConfigSetCmd(t *testing.T) {
// 	tests := []testutil.Test{
// 		{
// 			Name:     "Valid Arg for Editor",
// 			Args:     []string{"config", "set", "editor"},
// 			Expected: "",
// 			Error:    false,
// 		},
// 		{
// 			Name:     "Invalid Key",
// 			Args:     []string{"config", "set", "sandwich"},
// 			Expected: `Error: "sandwich" is not a valid key`,
// 			Error:    true,
// 		},
// 	}
//
// 	testutil.TestCmd(t, testutil.SetupFunc(func() error {
// 		return nil
// 	}), rootCmd, tests)
// }
