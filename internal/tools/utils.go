package tools

import (
	"fmt"
	"os"
	"path"
)

// exported functions
var (
	IsValidHook               = isValidHook
	GitExists                 = gitExists
	HuskyExists               = huskyExists
	GetHuskyHooksDir          = getHuskyHooksDir
	GetGitHooksDir            = getGitHooksDir
	IsCI                      = isCI
	ValidHooks                = validHooks
	ValidHooksWithDescription = validHooksWithDescription
)

const (
	green = "\033[0;32m"
	nc    = "\033[0m"
)

// internal functions
var validHooksWithDescription = []string{
	// Hooks de commit
	fmt.Sprintf("%s[pre-commit]%s             Execute before commit", green, nc),
	fmt.Sprintf("%s[prepare-commit-msg]%s     Execute before commit-msg", green, nc),
	fmt.Sprintf("%s[commit-msg]%s             Execute before commit", green, nc),
	fmt.Sprintf("%s[post-commit]%s            Execute after commit", green, nc),
	fmt.Sprintf("%s[post-commit-msg]%s        Execute after commit-msg\n", green, nc),
	fmt.Sprintf("%s[pre-merge]%s              Execute before merge", green, nc),
	fmt.Sprintf("%s[pre-merge-commit]%s       Execute before merge-commit", green, nc),
	fmt.Sprintf("%s[post-merge]%s             Execute after merge", green, nc),
	fmt.Sprintf("%s[post-merge-commit]%s      Execute after merge-commit\n", green, nc),
	fmt.Sprintf("%s[pre-rebase]%s             Execute before rebase", green, nc),
	fmt.Sprintf("%s[pre-rebase-commit]%s      Execute before rebase-commit", green, nc),
	fmt.Sprintf("%s[post-rebase]%s            Execute after rebase", green, nc),
	fmt.Sprintf("%s[post-rebase-commit]%s     Execute after rebase-commit\n", green, nc),
	fmt.Sprintf("%s[pre-push]%s               Execute before push", green, nc),
	fmt.Sprintf("%s[update]%s                 Execute after push\n", green, nc),
	fmt.Sprintf("%s[pre-applypatch]%s         Execute before applypatch", green, nc),
	fmt.Sprintf("%s[post-applypatch]%s        Execute after applypatch\n", green, nc),
	fmt.Sprintf("%s[post-checkout]%s          Execute after checkout", green, nc),
}

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
func isValidHook(hook string) bool {
	for _, h := range validHooks {
		if h == hook {
			return true
		}
	}
	return false
}

// GitExists checks if .git is installed
func gitExists() bool {
	_, err := os.Stat(".git")
	return err == nil
}

// HuskyExists checks if .husky is installed
func huskyExists() bool {
	_, err := os.Stat(".husky")
	return err == nil
}

// GetHuskyHooksDir returns the path to the husky hooks directory
func getHuskyHooksDir(relative bool) string {
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
func getGitHooksDir(relative bool) string {
	if relative {
		return path.Join(".git", "hooks")
	}
	cwd, err := os.Getwd()
	if err != nil {
		return ""
	}
	return path.Join(cwd, ".git", "hooks")
}

// IsCI checks if the current environment is a CI environment
func isCI() bool {
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

const HuskyGolang = `
  _    _                 _                 _____           _                         
 | |  | |               | |               / ____|         | |                        
 | |__| |  _   _   ___  | | __  _   _    | |  __    ___   | |   __ _   _ __     __ _ 
 |  __  | | | | | / __| | |/ / | | | |   | | |_ |  / _ \  | |  / _  | | '_ \   / _  |
 | |  | | | |_| | \__ \ |   <  | |_| |   | |__| | | (_) | | | | (_| | | | | | | (_| |
 |_|  |_|  \__,_| |___/ |_|\_\  \__, |    \_____|  \___/  |_|  \__,_| |_| |_|  \__, |
                                 __/ |                                          __/ |
                                |___/                                          |___/ 

`
