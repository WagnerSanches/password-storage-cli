package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// Encrypt recebe os dados em texto puro e a chave de 32 bytes do Argon2id
func Encrypt(plaintext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// GCM é o modo que garante integridade (autenticação)
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// O Nonce (Number used once) deve ser único para cada criptografia
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// Retornamos o nonce colado na frente do texto cifrado
	// O nonce não é secreto, mas é necessário para decriptar
	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

// Decrypt recebe o dado cifrado e a chave para retornar o texto original
func Decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext muito curto")
	}

	// Separamos o nonce (prefixo) do conteúdo cifrado real
	nonce, actualCiphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, actualCiphertext, nil)
}