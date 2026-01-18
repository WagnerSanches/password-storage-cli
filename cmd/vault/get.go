package main

import (
	"fmt"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/term"
	"github.com/WagnerSanches/vault/internal/crypto"
	"github.com/WagnerSanches/vault/internal/storage"
)

var getCmd = &cobra.Command{
	Use:   "get [serviço]",
	Short: "Recupera e exibe a senha de um serviço",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		service := args[0]

		data, err := storage.LoadFile("vault.json")
		if err != nil || len(data.Entries) == 0 {
			fmt.Println("Cofre vazio ou não encontrado.")
			return
		}

		var entry storage.Entry
		found := false
		for _, e := range data.Entries {
			if e.Service == service {
				entry = e
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("Serviço '%s' não encontrado.\n", service)
			return
		}

		fmt.Print("Digite a Master Password: ")
		masterPwd, _ := term.ReadPassword(int(syscall.Stdin))
		fmt.Println()

		key := crypto.DeriveKey(string(masterPwd), entry.Salt)

		decrypted, err := crypto.Decrypt(entry.Ciphertext, key)
		if err != nil {
			fmt.Println("Master Password incorreta.")
			return
		}

		fmt.Printf("Senha para '%s': %s\n", service, string(decrypted))
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
