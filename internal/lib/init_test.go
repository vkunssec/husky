package lib

import (
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	// Setup
	tmpDir := t.TempDir()
	originalWd, _ := os.Getwd()
	os.Chdir(tmpDir)
	defer os.Chdir(originalWd)

	// Create .git directory to simulate repository
	os.Mkdir(".git", 0755)

	tests := []struct {
		name    string
		opts    InitOptions
		wantErr bool
	}{
		{
			name: "Successful initialization",
			opts: InitOptions{
				Config:    NewDefaultConfig(),
				Templates: map[string]*HookTemplate{},
				Force:     false,
			},
			wantErr: false,
		},
		{
			name: "Forced initialization",
			opts: InitOptions{
				Config:    NewDefaultConfig(),
				Templates: map[string]*HookTemplate{},
				Force:     true,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Init(tt.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("Init() erro = %v, wantErr %v", err, tt.wantErr)
			}

			// Verify if .husky directory was created
			if _, err := os.Stat(".husky"); os.IsNotExist(err) {
				t.Error(".husky directory was not created")
			}
		})
	}
}
