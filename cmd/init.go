package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vkunssec/husky/internal/lib"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize husky",
	Long:  "Initialize husky in the current directory",
	Run: func(cmd *cobra.Command, args []string) {
		err := lib.Init()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
