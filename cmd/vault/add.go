package main

import (
	"fmt"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/term"
	"github.com/wagnersanches/vault/internal/crypto"
	"github.com/wagnersanches/vault/internal/storage"
)

var addCmd = &cobra.Command{
	Use:   "add [serviço]",
	Short: "Adiciona uma nova credencial ao cofre",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		service := args[0]

		fmt.Print("Digite a Master Password: ")
		masterPwd, _ := term.ReadPassword(int(syscall.Stdin))
		fmt.Println()

		fmt.Print("Digite a Senha do Serviço: ")
		servicePwd, _ := term.ReadPassword(int(syscall.Stdin))
		fmt.Println()

		salt := make([]byte, 16)

		key := crypto.DeriveKey(string(masterPwd), salt)

		ciphertext, _ := crypto.Encrypt(servicePwd, key)

		data, _ := storage.LoadFile("vault.json")
		newEntry := storage.Entry{
			Service:    service,
			Ciphertext: ciphertext,
			Salt:       salt,
		}
		data.Entries = append(data.Entries, newEntry)
		storage.SaveFile("vault.json", data)

		fmt.Printf("Senha '%s' salva com sucesso!\n", service)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}