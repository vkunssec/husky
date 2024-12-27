package lib

import (
	"os"
	"path"
	"testing"
)

func TestAdd(t *testing.T) {
	// Setup
	tmpDir := t.TempDir()
	originalWd, _ := os.Getwd()
	os.Chdir(tmpDir)
	defer os.Chdir(originalWd)

	// Criar estrutura necessária
	os.Mkdir(".git", 0755)
	os.MkdirAll(".husky/hooks", 0755)

	tests := []struct {
		name    string
		hook    string
		cmd     string
		wantErr bool
	}{
		{
			name:    "Valid hook",
			hook:    "pre-commit",
			cmd:     "echo 'teste'",
			wantErr: false,
		},
		{
			name:    "Invalid hook",
			hook:    "hook-invalido",
			cmd:     "echo 'teste'",
			wantErr: true,
		},
		{
			name:    "Empty command",
			hook:    "pre-commit",
			cmd:     "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Add(tt.hook, tt.cmd)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() erro = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				// Verify if the hook file was created
				hookPath := path.Join(".husky/hooks", tt.hook)
				if _, err := os.Stat(hookPath); os.IsNotExist(err) {
					t.Errorf("Hook file was not created in %s", hookPath)
				}
			}
		})
	}
}
