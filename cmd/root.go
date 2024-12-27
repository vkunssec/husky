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
		Long: `Husky é um gerenciador de Git hooks que permite configurar e gerenciar 
seus hooks de forma simples e eficiente.

Características principais:
- Configuração via arquivo yaml/json
- Suporte a múltiplos hooks
- Fácil instalação e uso
- Compatível com todos os sistemas operacionais

Para mais informações visite: https://github.com/vkunssec/husky`,
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
