package lib

import (
	"os"
	"path"
)

var validHooks = []string{
	"pre-commit",
	"pre-push",
	"pre-merge",
	"pre-rebase",
	"pre-merge-commit",
	"pre-rebase-commit",
	"pre-commit-msg",
	"post-commit",
	"post-merge",
	"post-rebase",
	"post-merge-commit",
	"post-rebase-commit",
	"update",
	"prepare-commit-msg",
	"commit-msg",
	"post-checkout",
	"pre-applypatch",
	"post-applypatch",
	"post-commit-msg",
}

// IsValidHook checks if the hook is valid
func IsValidHook(hook string) bool {
	for _, h := range validHooks {
		if h == hook {
			return true
		}
	}
	return false
}

// GitExists checks if .git is installed
func GitExists() bool {
	_, err := os.Stat(".git")
	return err == nil
}

// HuskyExists checks if .husky is installed
func HuskyExists() bool {
	_, err := os.Stat(".husky")
	return err == nil
}

// GetHuskyHooksDir returns the path to the husky hooks directory
func GetHuskyHooksDir(relative bool) string {
	if relative {
		return path.Join(".husky", "hooks")
	}
	cwd, err := os.Getwd()
	if err != nil {
		return ""
	}
	return path.Join(cwd, ".husky", "hooks")
}

// GetGitHooksDir returns the path to the git hooks directory
func GetGitHooksDir(relative bool) string {
	if relative {
		return path.Join(".git", "hooks")
	}
	cwd, err := os.Getwd()
	if err != nil {
		return ""
	}
	return path.Join(cwd, ".git", "hooks")
}
