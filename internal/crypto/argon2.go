package crypto

import (
	"golang.org/x/crypto/argon2"
)

const (
	Memory 		= 1024 * 64 // 64 mb
	Iterations 	= 3
	Parallelism = 2
	KeyLength 	= 32
)

func DeriveKey(password string, salt []byte) []byte {
	return argon2.IDKey(
		[]byte(password),
		salt,
		Iterations,
		Memory,
		Parallelism,
		KeyLength,
	)
}