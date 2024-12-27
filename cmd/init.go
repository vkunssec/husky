package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vkunssec/husky/internal/lib"
)

var (
	quiet bool
	force bool
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize husky",
	Long: `Initialize husky in the current directory.
	
Este comando irá:
- Criar o arquivo de configuração husky.yaml
- Configurar a estrutura básica de hooks
- Preparar o ambiente git`,
	Run: func(cmd *cobra.Command, args []string) {
		if !quiet {
			lib.LogInfo("Initializing Husky...")
		}

		opts := lib.InitOptions{
			Config:    lib.NewDefaultConfig(),
			Templates: lib.LoadTemplates(),
			Force:     force,
		}

		if err := lib.Init(opts); err != nil {
			lib.LogError("❌ Error initializing Husky: %v\n", err)
			return
		}

		if !quiet {
			lib.LogInfo("✅ Husky initialized successfully!")
		}
	},
}

func init() {
	initCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Modo silencioso")
	initCmd.Flags().BoolVarP(&force, "force", "f", false, "Force initialization")
	rootCmd.AddCommand(initCmd)
}
