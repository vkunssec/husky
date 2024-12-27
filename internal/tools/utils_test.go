package tools

import (
	"os"
	"testing"
)

func TestIsValidHook(t *testing.T) {
	tests := []struct {
		name string
		hook string
		want bool
	}{
		{"Valid pre-commit hook", "pre-commit", true},
		{"Valid post-commit hook", "post-commit", true},
		{"Invalid hook", "invalid-hook", false},
		{"Empty hook", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidHook(tt.hook); got != tt.want {
				t.Errorf("IsValidHook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGitExists(t *testing.T) {
	// Setup
	tmpDir := t.TempDir()
	originalWd, _ := os.Getwd()
	os.Chdir(tmpDir)
	defer os.Chdir(originalWd)

	t.Run("Git does not exist", func(t *testing.T) {
		if GitExists() {
			t.Error("GitExists() returned true when it should not exist .git")
		}
	})

	t.Run("Git exists", func(t *testing.T) {
		os.Mkdir(".git", 0755)
		if !GitExists() {
			t.Error("GitExists() returned false when it should exist .git")
		}
	})
}

func TestIsCI(t *testing.T) {
	tests := []struct {
		name    string
		envVar  string
		envVal  string
		wantCI  bool
		cleanup bool
	}{
		{"Generic CI environment", "CI", "true", true, true},
		{"GitHub Actions", "GITHUB_ACTIONS", "true", true, true},
		{"No CI", "", "", false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envVar != "" {
				os.Setenv(tt.envVar, tt.envVal)
			}
			if tt.cleanup {
				defer os.Unsetenv(tt.envVar)
			}

			if got := IsCI(); got != tt.wantCI {
				t.Errorf("IsCI() = %v, want %v", got, tt.wantCI)
			}
		})
	}
}
