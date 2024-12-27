package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootCmd(t *testing.T) {
	t.Run("should have correct version", func(t *testing.T) {
		assert.Equal(t, "1.0.2", version)
	})

	t.Run("should have correct command properties", func(t *testing.T) {
		assert.Equal(t, "husky", rootCmd.Use)
		assert.Equal(t, "Git hooks manager", rootCmd.Short)
		assert.Contains(t, rootCmd.Long, "Husky is a Git hooks manager")
	})
}

func TestExecute(t *testing.T) {
	t.Run("should execute without error", func(t *testing.T) {
		// Captura a saída do comando
		oldOut := rootCmd.OutOrStdout()
		defer func() { rootCmd.SetOut(oldOut) }()

		buf := new(bytes.Buffer)
		rootCmd.SetOut(buf)

		// Executa o comando
		Execute()

		// Verifica se não houve erro (se houvesse, o programa teria terminado com os.Exit(1))
		assert.NotPanics(t, func() { Execute() })
	})
}

func TestInit(t *testing.T) {
	t.Run("should have toggle flag", func(t *testing.T) {
		flag := rootCmd.PersistentFlags().Lookup("toggle")
		assert.NotNil(t, flag)
		assert.Equal(t, "toggle", flag.Name)
		assert.Equal(t, "t", flag.Shorthand)
		assert.Equal(t, "false", flag.DefValue)
	})
}
