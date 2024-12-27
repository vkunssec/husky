package lib

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
)

// Add adds a hook to the husky hooks directory
func Add(hook string, cmd string) error {
	if !IsValidHook(hook) {
		return errors.New("invalid hook")
	}

	// check if .git exists
	if !GitExists() {
		return errors.New("git not initialized")
	}

	// check if .husky exists
	if !HuskyExists() {
		return errors.New(".husky not initialized")
	}

	// check if .husky/hooks exists
	_, err := os.Stat(GetHuskyHooksDir(true))
	if os.IsNotExist(err) {
		LogInfo("no pre-existing hooks found")

		// create .husky/hooks
		err = os.MkdirAll(GetHuskyHooksDir(true), 0755)
		if err != nil {
			return err
		}

		LogInfo("created .husky/hooks")
	}

	// check if hook already exists
	if _, err := os.Stat(path.Join(GetHuskyHooksDir(true), hook)); err == nil {
		// ask if user wants to overwrite
		fmt.Printf("Hook '%s' already exists. Do you want to overwrite it? [y/N] ", hook)
		var response string
		fmt.Scanln(&response)
		if response != "y" && response != "Y" {
			return fmt.Errorf("operation cancelled by user")
		}
	}

	// create hook
	file, err := os.Create(path.Join(GetHuskyHooksDir(true), hook))
	if err != nil {
		return err
	}

	defer file.Close()

	if cmd == "" {
		return errors.New("command cannot be empty")
	}

	// Add shebang only if it doesn't exist
	if !strings.HasPrefix(cmd, "#!/") {
		cmd = "#!/bin/sh\n" + cmd
	}

	_, err = file.WriteString(cmd)
	if err != nil {
		return err
	}

	// Add execution permission to the file
	if err := os.Chmod(path.Join(GetHuskyHooksDir(true), hook), 0755); err != nil {
		return fmt.Errorf("failed to set hook permissions: %w", err)
	}

	return nil
}
