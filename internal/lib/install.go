package lib

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Install installs husky
func Install() error {
	fmt.Println("Installing husky")

	if !GitExists() {
		return errors.New("git is not installed")
	}

	if !HuskyExists() {
		return errors.New("husky is already installed")
	}

	gitHooksDir := GetGitHooksDir(true)
	huskyHooksDir := GetHuskyHooksDir(true)

	_, err := os.Stat(huskyHooksDir)
	if err != nil {
		return err
	}

	if err := os.RemoveAll(gitHooksDir); err != nil {
		return err
	}

	if err := os.Mkdir(gitHooksDir, 0755); err != nil {
		return err
	}

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

	for _, hook := range hooks {
		if hook == huskyHooksDir {
			continue
		}

		fmt.Print(hook)

		err = os.Link(hook, filepath.Join(gitHooksDir, filepath.Base(hook)))
		if err != nil {
			return err
		}

		err = os.Chmod(filepath.Join(gitHooksDir, filepath.Base(hook)), 0755)
		if err != nil {
			return err
		}
	}
	fmt.Println("Hooks installed")

	return nil
}
