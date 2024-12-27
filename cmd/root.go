package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/vkunssec/husky/internal/lib"
)

var (
	version = "0.0.4"

	rootCmd = &cobra.Command{
		Use:     "husky",
		Version: version,
		Short:   "Git hooks manager",
		Long: `Husky is a Git hooks manager that allows you to configure and manage 
your hooks in a simple and efficient way.

Main features:
- Configuration via yaml/json file
- Support for multiple hooks
- Easy installation and usage
- Compatible with all operating systems

For more information visit: https://github.com/vkunssec/husky`,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		lib.LogError("❌ Error executing command: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("toggle", "t", false, "Help message for toggle")
}