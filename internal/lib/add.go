package lib

import (
	"errors"
	"fmt"
	"os"
	"path"
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
		fmt.Println("no pre-existing hooks found")

		// create .husky/hooks
		err = os.MkdirAll(GetHuskyHooksDir(true), 0755)
		if err != nil {
			return err
		}

		fmt.Println("created .husky/hooks")
	}

	// create hook
	file, err := os.Create(path.Join(GetHuskyHooksDir(true), hook))
	if err != nil {
		return err
	}

	//goland:noinspection GoUnhandledErrorResult
	defer file.Close()

	cmd = "#!/bin/sh\n" + cmd
	_, err = file.WriteString(cmd)
	if err != nil {
		return err
	}

	return nil
}
