package tools

import (
	"os"
	"path"
)

var validHooks = []string{
	// Hooks de commit
	"pre-commit",
	"prepare-commit-msg",
	"commit-msg",
	"post-commit",
	"post-commit-msg",

	// Hooks de merge
	"pre-merge",
	"pre-merge-commit",
	"post-merge",
	"post-merge-commit",

	// Hooks de rebase
	"pre-rebase",
	"pre-rebase-commit",
	"post-rebase",
	"post-rebase-commit",

	// Hooks de push
	"pre-push",
	"update",

	// Hooks de patch
	"pre-applypatch",
	"post-applypatch",

	// Outros hooks
	"post-checkout",
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

// IsCI verifica se está rodando em um ambiente de CI
func IsCI() bool {
	ciEnvVars := []string{
		"CI",
		"TRAVIS",
		"CIRCLECI",
		"GITHUB_ACTIONS",
		"GITLAB_CI",
		"JENKINS_URL",
	}

	for _, env := range ciEnvVars {
		if os.Getenv(env) != "" {
			return true
		}
	}
	return false
}
