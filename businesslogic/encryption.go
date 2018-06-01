package businesslogic

import (
	crand "crypto/rand"
	"crypto/sha256"
	"io"
)

const (
	PW_SALT_BYTES = 32
	//CHARS = [] rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
)

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
