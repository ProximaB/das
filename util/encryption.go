package util

import (
	crand "crypto/rand"
	"crypto/sha256"
	"golang.org/x/crypto/bcrypt"
	"io"
)

const (
	PW_SALT_BYTES = 32
)

type IEncryptionStrategy interface {
	GenerateSalt(password string) []byte
	GenerateHash(password string) []byte
}

type Sha256EncryptionStrategy struct {
	Salt []byte
}

type BcryptEncryptionStrategy struct {
	Cost int
}

func GenerateSalt(password []byte) []byte {
	buff := make([]byte, PW_SALT_BYTES, PW_SALT_BYTES+sha256.Size)
	_, err := io.ReadFull(crand.Reader, buff)
	if err != nil {
		// TODO: how to handle this error?
	}
	hash := sha256.New()
	hash.Write(buff)
	hash.Write(password)
	return hash.Sum(buff)
}

func GenerateHash(salt []byte, password []byte) []byte {
	combination := string(salt) + string(password)
	passwordHash := sha256.New()
	io.WriteString(passwordHash, combination)
	return passwordHash.Sum(nil)
}

func (strategy BcryptEncryptionStrategy) GenerateSalt(password string) []byte {
	return nil
}

func (strategy BcryptEncryptionStrategy) GenerateHash(password string) []byte {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), strategy.Cost)
	if err != nil {
		return nil
	}
	return bytes
}
