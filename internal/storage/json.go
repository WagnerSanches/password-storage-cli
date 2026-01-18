package storage

import (
	"encoding/json"
	"os"
)

type Entry struct {
	Service    string `json:"service"`
	Username   string `json:"username"`
	Ciphertext []byte `json:"ciphertext"`
	Salt       []byte `json:"salt"`
	Nonce      []byte `json:"nonce"`
}

type VaultData struct {
	Entries []Entry `json:"entries"`
}

func SaveFile(path string, data VaultData) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, jsonData, 0600)
}

func LoadFile(path string) (VaultData, error) {
	var data VaultData

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return data, nil 
	}

	fileData, err := os.ReadFile(path)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(fileData, &data)
	return data, err
}