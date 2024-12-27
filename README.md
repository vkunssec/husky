# Husky

Husky is a Git hook manager that allows you to configure and manage your hooks in a simple and efficient way.

## Features

- Support for multiple hooks
- Simple installation and usage
- Compatible with all operating systems

## Installation

```bash
go install github.com/vkunssec/husky@latest
```

## Usage

### Initialization

To start using Husky in your project:

```bash
husky init
```

To force initialization in an existing project:

```bash
husky init --force
```

This command will:
- Configure the basic hook structure
- Prepare the Git environment

### Adding Hooks

To add a new hook:

```bash
husky add <hook-name> '<comando>'
```

Example:
```bash
husky add pre-commit 'go test ./...'
```

### Supported Hooks

- Commit Hooks
  - pre-commit
  - prepare-commit-msg
  - commit-msg
  - post-commit
  - post-commit-msg

- Merge Hooks
  - pre-merge
  - pre-merge-commit
  - post-merge
  - post-merge-commit

- Rebase Hooks
  - pre-rebase
  - pre-rebase-commit
  - post-rebase
  - post-rebase-commit

- Push Hooks
  - pre-push
  - update

- Patch Hooks
  - pre-applypatch
  - post-applypatch

- Others
  - post-checkout

### Installing Hooks

To install the configured hooks:

```bash
husky install   
```

## Directory Structure

After initialization, Husky creates the following structure:

```
.
├── .git/
│   └── hooks/          # Git hooks (managed by Husky)
└── .husky/
    └── hooks/          # Your custom hooks
```

## Contributing

Contributions are welcome! Please feel free to submit pull requests.

## License

[MIT License](LICENSE)

## Author

[Victor Martins](https://github.com/vkunssec)
