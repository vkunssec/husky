package lib

import (
	"errors"
	"fmt"
	"os"
	"path"
	"time"
)

// InitOptions are the options for the init command
type InitOptions struct {
	Config    *HuskyConfig             // Husky configuration
	Templates map[string]*HookTemplate // Hook templates
	Force     bool                     // Force initialization
}

// Init initializes husky
func Init(opts InitOptions) error {
	// Validate environment
	if err := validateEnvironment(!opts.Force); err != nil {
		return fmt.Errorf("environment validation failed: %w", err)
	}

	// Create backup if enabled
	if opts.Config.BackupEnabled {
		if err := createBackup(); err != nil {
			LogError("Failed to create backup: %v", err)
		}
	}

	// Create husky directory structure
	huskyDir, err := createHuskyStructure(opts.Config)
	if err != nil {
		return fmt.Errorf("failed to create husky structure: %w", err)
	}

	// Install default hooks
	if err := installDefaultHooks(huskyDir, opts); err != nil {
		cleanup(huskyDir)
		return fmt.Errorf("failed to install default hooks: %w", err)
	}

	LogInfo("Husky initialized successfully")
	return nil
}

// validateEnvironment validates the environment
func validateEnvironment(checkExisting bool) error {
	if !GitExists() {
		return errors.New("git repository not initialized")
	}

	if checkExisting && HuskyExists() {
		return errors.New("husky already initialized")
	}

	return nil
}

// createHuskyStructure creates the husky directory structure
func createHuskyStructure(config *HuskyConfig) (string, error) {
	huskyDir := GetHuskyHooksDir(true)

	if err := os.MkdirAll(huskyDir, config.DefaultPermissions); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}

	return huskyDir, nil
}

// installDefaultHooks installs the default hooks
func installDefaultHooks(huskyDir string, opts InitOptions) error {
	for hookName, template := range opts.Config.DefaultHooks {
		if err := createHook(huskyDir, hookName, template, opts.Config); err != nil {
			return fmt.Errorf("failed to create %s hook: %w", hookName, err)
		}
	}
	return nil
}

// createHook creates a hook
func createHook(dir, name, content string, config *HuskyConfig) error {
	hookPath := path.Join(dir, name)

	file, err := os.OpenFile(hookPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, config.DefaultPermissions)
	if err != nil {
		return err
	}

	defer func() {
		if cerr := file.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	if _, err = file.WriteString(content); err != nil {
		return err
	}

	return nil
}

// cleanup cleans up the husky directory
func cleanup(dir string) {
	if err := os.RemoveAll(dir); err != nil {
		LogError("Failed to cleanup directory: %v", err)
	}
}

// createBackup creates a backup of the current git hooks
func createBackup() error {
	// Get source and destination directories
	sourceDir := GetGitHooksDir(true)
	backupDir := path.Join(GetHuskyHooksDir(true), "backup", fmt.Sprintf("hooks_%d", time.Now().Unix()))

	// Create backup directory
	if err := os.MkdirAll(backupDir, 0755); err != nil {
		return fmt.Errorf("failed to create backup directory: %w", err)
	}

	// Read all files from the hooks directory
	entries, err := os.ReadDir(sourceDir)
	if err != nil {
		return fmt.Errorf("failed to read hooks directory: %w", err)
	}

	// Copy each hook to the backup directory
	for _, entry := range entries {
		if entry.IsDir() {
			continue // Skip directories
		}

		sourcePath := path.Join(sourceDir, entry.Name())
		destPath := path.Join(backupDir, entry.Name())

		// Read source file
		content, err := os.ReadFile(sourcePath)
		if err != nil {
			return fmt.Errorf("failed to read hook %s: %w", entry.Name(), err)
		}

		// Write to backup file
		if err := os.WriteFile(destPath, content, 0755); err != nil {
			return fmt.Errorf("failed to write backup of hook %s: %w", entry.Name(), err)
		}
	}

	LogInfo("Backup of hooks created in: %s", backupDir)
	return nil
}
