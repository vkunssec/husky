package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vkunssec/husky/internal/lib"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install husky",
	Long: `Install husky in the current directory.
	
Este comando irá:
- Verificar pré-requisitos
- Instalar os hooks configurados
- Configurar os scripts git`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := lib.Install(); err != nil {
			lib.LogError("❌ Error installing Husky: %v\n", err)
			return
		}

		lib.LogInfo("✅ Husky installed successfully!")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
