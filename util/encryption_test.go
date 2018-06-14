package util_test

import (
	"github.com/DancesportSoftware/das/util"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestBcryptEncryptionStrategy_GenerateHash(t *testing.T) {
	strategy := util.BcryptEncryptionStrategy{Cost: 10}

	pw1 := "d*3lH\130Jz#e"
	pw2 := "4bv!pYn8f6dS$"

	hash_1 := strategy.GenerateHash(pw1)
	hash_2 := strategy.GenerateHash(pw2)

	assert.False(t, reflect.DeepEqual(hash_1, hash_2), "different password should result in unequal hashes")
}
