package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vkunssec/husky/internal/lib"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install husky",
	Long:  "Install husky in the current directory",
	Run: func(cmd *cobra.Command, args []string) {
		err := lib.Install()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
