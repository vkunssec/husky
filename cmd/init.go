package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vkunssec/husky/internal/lib"
	"github.com/vkunssec/husky/internal/tools"
)

var (
	quiet bool
	force bool
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize husky",
	Long: `Initialize husky in the current directory.
	
This command will:
- Configure the basic hook structure
- Prepare the git environment`,
	Run: func(cmd *cobra.Command, args []string) {
		if !quiet {
			tools.LogInfo("Initializing Husky...")
		}

		opts := lib.InitOptions{
			Config:    lib.NewDefaultConfig(),
			Templates: lib.LoadTemplates(),
			Force:     force,
			Quiet:     quiet,
		}

		optsInstall := lib.InstallOptions{
			Quiet: quiet,
		}

		if err := lib.Init(opts); err != nil {
			tools.LogError("❌ Error initializing Husky: %v\n", err)
			return
		}

		if err := lib.Install(optsInstall); err != nil {
			tools.LogError("❌ Error installing hooks: %v\n", err)
			return
		}

		if !quiet {
			tools.LogInfo("✅ Husky initialized successfully!")
		}
	},
}

func init() {
	initCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Silent mode")
	initCmd.Flags().BoolVarP(&force, "force", "f", false, "Force initialization")
	rootCmd.AddCommand(initCmd)
}
