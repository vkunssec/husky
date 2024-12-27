package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vkunssec/husky/internal/lib"
)

var addCmd = &cobra.Command{
	Use:     "add [hook] [comando]",
	Short:   "Add a hook",
	Long:    "Adiciona um novo hook ao repositório",
	Args:    cobra.ExactArgs(2),
	Example: "husky add pre-commit 'go test ./...'",
	Run: func(cmd *cobra.Command, args []string) {
		hook := args[0]
		cmdStr := args[1]

		if err := lib.Add(hook, cmdStr); err != nil {
			lib.LogError("❌ Error adding hook: %v\n", err)
			return
		}

		if err := lib.Install(); err != nil {
			lib.LogError("❌ Error installing hooks: %v\n", err)
			return
		}

		lib.LogInfo("✅ Hook '%s' added successfully!\n", hook)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
