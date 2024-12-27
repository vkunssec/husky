package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vkunssec/husky/internal/lib"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add a hook",
	Long:    "Add a hook to the repository",
	Args:    cobra.ExactArgs(2),
	Example: "husky add pre-commit 'echo \"husky installed\"'",
	Run: func(cmd *cobra.Command, args []string) {
		hook := args[0]
		cmdStr := args[1]

		err := lib.Add(hook, cmdStr)
		if err != nil {
			fmt.Println(err)
		}

		err = lib.Install()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
