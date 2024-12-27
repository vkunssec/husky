package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootCmd(t *testing.T) {
	t.Run("should have correct version", func(t *testing.T) {
		assert.Equal(t, "1.1.0", version)
	})

	t.Run("should have correct command properties", func(t *testing.T) {
		assert.Equal(t, "husky", rootCmd.Use)
		assert.Equal(t, "Git hooks manager", rootCmd.Short)
		assert.Contains(t, rootCmd.Long, "Husky is a Git hooks manager")
	})
}

func TestExecute(t *testing.T) {
	t.Run("should execute without error", func(t *testing.T) {
		// Capture command output
		oldOut := rootCmd.OutOrStdout()
		defer func() { rootCmd.SetOut(oldOut) }()

		buf := new(bytes.Buffer)
		rootCmd.SetOut(buf)

		// Execute command
		Execute()

		// Verify if there was no error (if there was, the program would have exited with os.Exit(1))
		assert.NotPanics(t, func() { Execute() })
	})
}
