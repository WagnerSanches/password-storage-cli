package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vault",
	Short: "Vault é um gerenciador de senhas seguro para seu Homelab",
	Long:  `Um CLI desenvolvido em Go que utiliza Argon2id e AES-256-GCM para proteger suas credenciais localmente.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}