package lib

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// Install installs husky git hooks by copying them from husky hooks directory to git hooks directory
func Install() error {
	LogInfo("Installing husky")

	// Check if git is installed in the system
	if !GitExists() {
		return errors.New("git is not installed")
	}

	// Check if husky is installed in the project
	if !HuskyExists() {
		return errors.New("husky is not installed")
	}

	// Get the paths for git hooks and husky hooks directories
	gitHooksDir := GetGitHooksDir(true)
	huskyHooksDir := GetHuskyHooksDir(true)

	// validate hook paths
	err := validateHookPath(huskyHooksDir, gitHooksDir)
	if err != nil {
		return err
	}

	// Verify if husky hooks directory exists
	_, err = os.Stat(huskyHooksDir)
	if err != nil {
		return err
	}

	// Remove existing git hooks directory if it exists
	if err := os.RemoveAll(gitHooksDir); err != nil {
		return err
	}

	// Create a new git hooks directory with proper permissions
	if err := os.Mkdir(gitHooksDir, 0700); err != nil {
		return err
	}

	// Store all hook paths from husky directory
	var hooks []string
	err = filepath.Walk(huskyHooksDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		hooks = append(hooks, path)
		return nil
	})
	if err != nil {
		return err
	}

	// Process each hook file
	for _, hook := range hooks {
		// Skip the husky hooks directory itself
		if hook == huskyHooksDir {
			continue
		}

		LogInfo(hook)

		// Create a hard link from husky hook to git hooks directory
		err = os.Link(hook, filepath.Join(gitHooksDir, filepath.Base(hook)))
		if err != nil {
			return err
		}

		// Set proper execution permissions for the hook
		err = os.Chmod(filepath.Join(gitHooksDir, filepath.Base(hook)), 0700)
		if err != nil {
			return err
		}
	}
	LogInfo("Hooks installed")

	return nil
}

// validateHookPath validates the hook path and base directory
func validateHookPath(hookPath string, baseDir string) error {
	// resolve hook path to absolute path
	absPath, err := filepath.Abs(hookPath)
	if err != nil {
		return err
	}

	// check if hook path is inside the base directory
	if !strings.HasPrefix(absPath, baseDir) {
		return errors.New("hook path outside of allowed directory")
	}

	// check if hook path is a symlink
	fi, err := os.Lstat(hookPath)
	if err != nil {
		return err
	}
	if fi.Mode()&os.ModeSymlink != 0 {
		return errors.New("symlinks not allowed")
	}

	return nil
}
