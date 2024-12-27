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
		opts := lib.InstallOptions{
			Quiet: quiet,
		}

		if err := lib.Install(opts); err != nil {
			tools.LogError("❌ Error installing Husky: %v\n", err)
			return
		}

		if !quiet {
			tools.LogInfo("✅ Husky installed successfully!")
		}
	},
}

func init() {
	installCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Silent mode")
	rootCmd.AddCommand(installCmd)
}
