// Dancesport Application System (DAS)
// Copyright (C) 2017, 2018 Yubing Hou
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

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
