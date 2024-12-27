package lib

import "os"

type HookTemplate struct {
	Name     string
	Content  string
	Validate func(string) error
}

const defaultPreCommitTemplate = `#!/bin/sh
# Husky pre-commit hook
set -e

# Validate that we're in a Git repository
if [ ! -d .git ]; then
    echo "Error: not a git repository"
    exit 1
fi

# Add your pre-commit commands here
`

const defaultPrePushTemplate = `#!/bin/sh
# Husky pre-push hook
set -e

# Add your pre-push commands here
`

const defaultPostCommitTemplate = `#!/bin/sh
# Husky post-commit hook
set -e

# Add your post-commit commands here
`

type HuskyConfig struct {
	DefaultPermissions os.FileMode
	HooksTemplatesDir  string
	DefaultHooks       map[string]string
	BackupEnabled      bool
	LogLevel           string
}

func NewDefaultConfig() *HuskyConfig {
	return &HuskyConfig{
		DefaultPermissions: 0755,
		HooksTemplatesDir:  "templates",
		DefaultHooks: map[string]string{
			"pre-commit":  defaultPreCommitTemplate,
			"pre-push":    defaultPrePushTemplate,
			"post-commit": defaultPostCommitTemplate,
		},
		BackupEnabled: true,
		LogLevel:      "info",
	}
}

func LoadTemplates() map[string]*HookTemplate {
	return map[string]*HookTemplate{
		"pre-commit":  {Name: "pre-commit", Content: defaultPreCommitTemplate},
		"pre-push":    {Name: "pre-push", Content: defaultPrePushTemplate},
		"post-commit": {Name: "post-commit", Content: defaultPostCommitTemplate},
	}
}
