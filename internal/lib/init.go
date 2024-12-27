package lib

import (
	"errors"
	"os"
	"path"
)

func Init() error {
	// check if .git exists
	if !GitExists() {
		return errors.New("git not initialized")
	}

	// check if .husky exists
	if HuskyExists() {
		return errors.New(".husky already exist")
	}

	// if not, create .husky/hooks
	err := os.MkdirAll(GetHuskyHooksDir(true), 0755)
	if err != nil {
		return err
	}

	// create default pre-commit hook
	file, err := os.Create(path.Join(GetHuskyHooksDir(true), "pre-commit"))
	if err != nil {
		return err
	}

	//goland:noinspection GoUnhandledErrorResult
	defer file.Close()

	_, err = file.WriteString(`#!/bin/sh\necho "husky installed"`)
	if err != nil {
		return err
	}

	// add hooks to .git/hooks
	err = Install()
	if err != nil {
		return err
	}

	return nil
}
