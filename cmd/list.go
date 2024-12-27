package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vkunssec/husky/internal/lib"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all hooks",
	Long:  "List all hooks implemented in the repository",
	Run: func(cmd *cobra.Command, args []string) {
		lib.List()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
