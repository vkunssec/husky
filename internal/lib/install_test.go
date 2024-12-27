package lib

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vkunssec/husky/internal/tools"
)

func TestInstall(t *testing.T) {
	// save original functions
	prevGitExists := tools.GitExists
	prevHuskyExists := tools.HuskyExists
	prevGetGitHooksDir := tools.GetGitHooksDir
	prevGetHuskyHooksDir := tools.GetHuskyHooksDir

	// restore original functions
	defer func() {
		tools.GitExists = prevGitExists
		tools.HuskyExists = prevHuskyExists
		tools.GetGitHooksDir = prevGetGitHooksDir
		tools.GetHuskyHooksDir = prevGetHuskyHooksDir
	}()

	// create temporary directory for tests
	tmpDir, err := os.MkdirTemp("", "husky-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// create necessary directories for tests
	gitHooksDir := filepath.Join(tmpDir, "git-hooks")
	huskyHooksDir := filepath.Join(tmpDir, "husky-hooks")

	tests := []struct {
		name    string
		setup   func()
		opts    InstallOptions
		wantErr bool
		errMsgs []string
	}{
		{
			name: "Successful installation",
			setup: func() {
				// configure mocks for success
				tools.GitExists = func() bool { return true }
				tools.HuskyExists = func() bool { return true }
				tools.GetGitHooksDir = func(abs bool) string { return gitHooksDir }
				tools.GetHuskyHooksDir = func(abs bool) string { return huskyHooksDir }

				// create husky directory with an example hook
				err := os.MkdirAll(huskyHooksDir, 0755)
				if err != nil {
					t.Fatal(err)
				}
				hookPath := filepath.Join(huskyHooksDir, "pre-commit")
				err = os.WriteFile(
					hookPath,
					[]byte("#!/bin/sh\necho 'test'"),
					0755,
				)
				if err != nil {
					t.Fatal(err)
				}
				// ensure correct permissions
				err = os.Chmod(hookPath, 0755)
				if err != nil {
					t.Fatal(err)
				}
			},
			opts:    InstallOptions{Quiet: false},
			wantErr: false,
		},
		{
			name: "Git not installed",
			setup: func() {
				tools.GitExists = func() bool { return false }
			},
			opts:    InstallOptions{Quiet: true},
			wantErr: true,
			errMsgs: []string{"git is not installed"},
		},
		{
			name: "Husky not installed",
			setup: func() {
				tools.GitExists = func() bool { return true }
				tools.HuskyExists = func() bool { return false }
			},
			opts:    InstallOptions{Quiet: true},
			wantErr: true,
			errMsgs: []string{"husky is not installed"},
		},
		{
			name: "Husky hooks directory not found",
			setup: func() {
				tools.GitExists = func() bool { return true }
				tools.HuskyExists = func() bool { return true }
				tools.GetGitHooksDir = func(abs bool) string { return gitHooksDir }
				tools.GetHuskyHooksDir = func(abs bool) string { return huskyHooksDir }

				// remove husky directory if it exists
				os.RemoveAll(huskyHooksDir)
			},
			opts:    InstallOptions{Quiet: true},
			wantErr: true,
			errMsgs: []string{
				"no such file or directory",
				"The system cannot find the file specified",
				"cannot find the path specified",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// clean state before each test
			os.RemoveAll(gitHooksDir)
			os.RemoveAll(huskyHooksDir)

			// create git hooks directory
			err := os.MkdirAll(gitHooksDir, 0755)
			if err != nil {
				t.Fatal(err)
			}

			// configure environment for test
			tt.setup()

			// execute test
			err = install(tt.opts)

			// check result
			if tt.wantErr {
				assert.Error(t, err)
				if len(tt.errMsgs) > 0 {
					errStr := err.Error()
					found := false
					for _, msg := range tt.errMsgs {
						if strings.Contains(errStr, msg) {
							found = true
							break
						}
					}
					assert.True(t, found,
						"Error '%s' does not contain any of the expected messages: %v",
						errStr, tt.errMsgs)
				}
			} else {
				assert.NoError(t, err)

				// check if hooks were installed correctly
				if !tt.wantErr {
					// check if git hooks directory exists
					_, err := os.Stat(gitHooksDir)
					assert.NoError(t, err)

					// check if hooks were copied
					files, err := os.ReadDir(huskyHooksDir)
					assert.NoError(t, err)

					for _, file := range files {
						gitHookPath := filepath.Join(gitHooksDir, file.Name())
						_, err := os.Stat(gitHookPath)
						assert.NoError(t, err)

						// check permissions
						info, err := os.Stat(gitHookPath)
						assert.NoError(t, err)

						// on Windows, permissions are different
						if runtime.GOOS == "windows" {
							// on Windows, we check only if the file is not read-only
							fileAttributes := info.Mode().Perm()
							isExecutable := fileAttributes&0111 != 0
							isWritable := fileAttributes&0222 != 0
							assert.True(t, isExecutable || isWritable,
								"The file must be executable or writable on Windows")
						} else {
							// on Unix-like systems, we check the exact permission
							assert.Equal(t, os.FileMode(0755), info.Mode()&0777,
								"Incorrect permissions on file")
						}
					}
				}
			}
		})
	}
}
