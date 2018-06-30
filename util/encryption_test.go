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
