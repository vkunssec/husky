package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vkunssec/husky/internal/lib"
	"github.com/vkunssec/husky/internal/tools"
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
			tools.LogError("❌ Error adding hook: %v\n", err)
			return
		}

		if err := lib.Install(lib.InstallOptions{Quiet: quiet}); err != nil {
			tools.LogError("❌ Error installing hooks: %v\n", err)
			return
		}

		tools.LogInfo("✅ Hook '%s' added successfully!\n", hook)
	},
}

func init() {
	addCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Silent mode")
	rootCmd.AddCommand(addCmd)
}
