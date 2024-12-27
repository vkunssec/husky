package cmd

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/vkunssec/husky/internal/lib"
)

// Define types for functions
type AddFunc func(hook, cmdStr string) error
type InstallFunc func(opts lib.InstallOptions) error

// Declare original variables with explicit types
var (
	originalAdd     = lib.Add
	originalInstall = lib.Install
)

func init() {
	// Initialize original functions
	originalAdd = lib.Add
	originalInstall = lib.Install
}

func TestAddCmd(t *testing.T) {
	// save original functions
	prevAdd := lib.Add
	prevInstall := lib.Install

	// Mock mais detalhado para Add
	lib.Add = func(hook, cmdStr string) error {
		// valid hooks
		validHooks := map[string]bool{
			"pre-commit":         true,
			"pre-push":           true,
			"post-merge":         true,
			"pre-rebase":         true,
			"post-checkout":      true,
			"commit-msg":         true,
			"prepare-commit-msg": true,
			"post-commit":        true,
			"pre-auto-gc":        true,
		}

		// explicit hook validation
		if _, exists := validHooks[hook]; !exists {
			return errors.New("invalid hook")
		}

		// empty command validation
		if cmdStr == "" {
			return errors.New("empty command")
		}

		return nil
	}

	lib.Install = func(opts lib.InstallOptions) error {
		return nil
	}

	defer func() {
		lib.Add = prevAdd
		lib.Install = prevInstall
	}()

	tests := []struct {
		name    string
		args    []string
		quiet   bool
		wantErr bool
		errMsg  string
	}{
		{
			name:    "Add hook with silent mode",
			args:    []string{"pre-push", "go vet ./...", "-q"},
			quiet:   true,
			wantErr: false,
		},
		{
			name:    "Error - Too many arguments",
			args:    []string{"pre-commit", "go test", "extra"},
			quiet:   false,
			wantErr: true,
			errMsg:  "accepts 2 arg(s), received 3",
		},
		{
			name:    "Error - Invalid hook",
			args:    []string{"invalid-hook", "go test"},
			quiet:   false,
			wantErr: true,
			errMsg:  "invalid hook",
		},
		{
			name:    "Add valid hook",
			args:    []string{"pre-commit", "go test ./..."},
			quiet:   false,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create a new command instance for each test
			cmd := &cobra.Command{
				Use:  "add [hook] [comando]",
				Args: cobra.ExactArgs(2),
				RunE: func(cmd *cobra.Command, args []string) error {
					hook := args[0]
					cmdStr := args[1]

					if err := lib.Add(hook, cmdStr); err != nil {
						return err
					}

					if err := lib.Install(lib.InstallOptions{Quiet: quiet}); err != nil {
						return err
					}

					return nil
				},
			}

			// configure buffer to capture output
			buf := new(bytes.Buffer)
			cmd.SetOut(buf)
			cmd.SetErr(buf)

			// configure flags
			cmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Silent mode")
			if tt.quiet {
				cmd.Flags().Set("quiet", "true")
			}

			// configure arguments
			cmd.SetArgs(tt.args)

			// execute command
			err := cmd.Execute()

			if tt.wantErr {
				assert.Error(t, err, "expected an error for case: %s", tt.name)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg,
						"expected error message not found for case: %s", tt.name)
				}
			} else {
				assert.NoError(t, err)
			}

			// check if the quiet flag was configured correctly
			quietFlag, _ := cmd.Flags().GetBool("quiet")
			assert.Equal(t, tt.quiet, quietFlag)
		})
	}
}

func TestAddCmdFlags(t *testing.T) {
	// check if the quiet flag was registered correctly
	assert.NotNil(t, addCmd.Flags().Lookup("quiet"))

	// check default value of the quiet flag
	quietFlag, err := addCmd.Flags().GetBool("quiet")
	assert.NoError(t, err)
	assert.False(t, quietFlag)
}

func TestAddCmdValidation(t *testing.T) {
	// save original functions
	prevAdd := lib.Add
	prevInstall := lib.Install

	// configure mocks
	lib.Add = func(hook, cmdStr string) error {
		// t.Logf("validating - hook: '%s', command: '%s'", hook, cmdStr)

		// basic validations
		if hook == "" || len(strings.TrimSpace(hook)) == 0 {
			// t.Log("empty hook detected")
			return errors.New("empty hook")
		}

		if cmdStr == "" || len(strings.TrimSpace(cmdStr)) == 0 {
			// t.Log("empty command detected")
			return errors.New("empty command")
		}

		// valid hooks
		validHooks := map[string]bool{
			"pre-commit":         true,
			"pre-push":           true,
			"post-merge":         true,
			"pre-rebase":         true,
			"post-checkout":      true,
			"commit-msg":         true,
			"prepare-commit-msg": true,
			"post-commit":        true,
			"pre-auto-gc":        true,
		}

		if _, exists := validHooks[hook]; !exists {
			// t.Logf("invalid hook detected: %s", hook)
			return errors.New("invalid hook")
		}

		return nil
	}

	lib.Install = func(opts lib.InstallOptions) error {
		return nil
	}

	// restore original functions after tests
	defer func() {
		lib.Add = prevAdd
		lib.Install = prevInstall
	}()

	tests := []struct {
		name    string
		args    []string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "Valid hook with simple command",
			args:    []string{"pre-commit", "go test"},
			wantErr: false,
		},
		{
			name:    "Valid hook with complex command",
			args:    []string{"pre-push", "go test ./... && go vet ./..."},
			wantErr: false,
		},
		{
			name:    "Missing arguments",
			args:    []string{},
			wantErr: true,
			errMsg:  "accepts 2 arg(s), received 0",
		},
		{
			name:    "Only hook argument",
			args:    []string{"pre-commit"},
			wantErr: true,
			errMsg:  "accepts 2 arg(s), received 1",
		},
		{
			name:    "Whitespace hook",
			args:    []string{"   ", "go test"},
			wantErr: true,
			errMsg:  "empty hook",
		},
		{
			name:    "Whitespace command",
			args:    []string{"pre-commit", "   "},
			wantErr: true,
			errMsg:  "empty command",
		},
		{
			name:    "Invalid hook",
			args:    []string{"invalid-hook", "go test"},
			wantErr: true,
			errMsg:  "invalid hook",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create a new command instance for each test
			cmd := &cobra.Command{
				Use:  "add [hook] [comando]",
				Args: cobra.ExactArgs(2),
				RunE: func(cmd *cobra.Command, args []string) error {
					hook := args[0]
					cmdStr := args[1]

					if err := lib.Add(hook, cmdStr); err != nil {
						return err
					}

					if err := lib.Install(lib.InstallOptions{Quiet: quiet}); err != nil {
						return err
					}

					return nil
				},
			}

			// configure buffer to capture output
			buf := new(bytes.Buffer)
			cmd.SetOut(buf)
			cmd.SetErr(buf)

			// configure flags
			cmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Silent mode")

			// configure arguments
			cmd.SetArgs(tt.args)

			// execute command
			err := cmd.Execute()

			if tt.wantErr {
				assert.Error(t, err, "expected an error for case: %s", tt.name)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg,
						"expected error message not found for case: %s", tt.name)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
