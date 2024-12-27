package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vkunssec/husky/internal/lib"
	"github.com/vkunssec/husky/internal/tools"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install husky",
	Long: `Install husky in the current directory.
	
This command will:
- Check prerequisites
- Install the configured hooks
- Configure the git scripts`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := lib.Install(); err != nil {
			tools.LogError("❌ Error installing Husky: %v\n", err)
			return
		}

		tools.LogInfo("✅ Husky installed successfully!")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
