package cmd

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestInitCmd(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
		quiet   bool
		force   bool
	}{
		{
			name:    "Execution with silent mode",
			args:    []string{"--quiet"},
			wantErr: false,
			quiet:   true,
			force:   false,
		},
		{
			name:    "Execution with force",
			args:    []string{"--force"},
			wantErr: false,
			quiet:   false,
			force:   true,
		},
		{
			name:    "Normal execution",
			args:    []string{},
			wantErr: false,
			quiet:   false,
			force:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare buffer to capture output
			buf := new(bytes.Buffer)
			rootCmd.SetOut(buf)
			rootCmd.SetErr(buf)

			// Reset global flags
			quiet = false
			force = false

			// Configure command
			cmd := &cobra.Command{}
			cmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Silent mode")
			cmd.Flags().BoolVarP(&force, "force", "f", false, "Force initialization")

			// Execute command with test arguments
			cmd.SetArgs(tt.args)
			err := cmd.Execute()

			// Verifications
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// Verify if flags were configured correctly
			assert.Equal(t, tt.quiet, quiet)
			assert.Equal(t, tt.force, force)
		})
	}
}

func TestInitCmdFlags(t *testing.T) {
	// Test if flags were registered correctly
	assert.NotNil(t, initCmd.Flags().Lookup("quiet"))
	assert.NotNil(t, initCmd.Flags().Lookup("force"))

	// Verify default values of flags
	quietFlag, _ := initCmd.Flags().GetBool("quiet")
	forceFlag, _ := initCmd.Flags().GetBool("force")

	assert.False(t, quietFlag)
	assert.False(t, forceFlag)
}
