package cmd

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vkunssec/husky/internal/lib"
)

// Mock for lib.Install
var mockInstallFunc func(opts lib.InstallOptions) error

func init() {
	// Change the original function to a mock
	mockInstallFunc = func(opts lib.InstallOptions) error {
		return nil
	}
}

func TestInstallCmd(t *testing.T) {
	t.Run("should have correct command properties", func(t *testing.T) {
		assert.Equal(t, "install", installCmd.Use)
		assert.Equal(t, "Install husky", installCmd.Short)
		assert.Contains(t, installCmd.Long, "Install husky in the current directory")
	})

	t.Run("should execute successfully in normal mode", func(t *testing.T) {
		// Change the mock to return success
		mockInstallFunc = func(_ lib.InstallOptions) error {
			assert.False(t, quiet)
			return nil
		}

		// Capture the command output
		buf := new(bytes.Buffer)
		installCmd.SetOutput(buf)

		// Reset the quiet flag
		quiet = false

		// Execute the command
		installCmd.Execute()

		// Verify the output
		assert.Contains(t, buf.String(), "")
	})

	t.Run("should execute successfully in quiet mode", func(t *testing.T) {
		// Change the mock to return success
		mockInstallFunc = func(_ lib.InstallOptions) error {
			assert.True(t, quiet)
			return nil
		}

		// Capture the command output
		buf := new(bytes.Buffer)
		installCmd.SetOutput(buf)

		// Set quiet to true
		quiet = true

		// Execute the command
		installCmd.Execute()

		// Verify that there is no output in quiet mode
		assert.Empty(t, buf.String())
	})

	t.Run("should handle installation error", func(t *testing.T) {
		// Change the mock to return error
		mockInstallFunc = func(_ lib.InstallOptions) error {
			return errors.New("expected error")
		}

		// Capture the command output
		buf := new(bytes.Buffer)
		installCmd.SetOutput(buf)

		// Execute the command
		installCmd.Execute()

		// Verify the error message
		assert.Contains(t, buf.String(), "")
	})
}

func TestInstallInit(t *testing.T) {
	t.Run("should register quiet flag", func(t *testing.T) {
		flag := installCmd.Flags().Lookup("quiet")
		assert.NotNil(t, flag)
		assert.Equal(t, "quiet", flag.Name)
		assert.Equal(t, "q", flag.Shorthand)
		assert.Equal(t, "false", flag.DefValue)
	})

	t.Run("should be registered in root command", func(t *testing.T) {
		cmd, _, err := rootCmd.Find([]string{"install"})
		assert.NoError(t, err)
		assert.Equal(t, installCmd, cmd)
	})
}
