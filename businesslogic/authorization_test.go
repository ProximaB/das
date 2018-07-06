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

package businesslogic_test

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/mock/businesslogic"
	"github.com/DancesportSoftware/das/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthenticateUser_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	username := "testuser"
	password := "testpassword"
	userRepo := mock_businesslogic.NewMockIAccountRepository(mockCtrl)

	salt := util.GenerateSalt([]byte(password))
	hash := util.GenerateHash(salt, []byte(password))

	// specifies behavior
	userRepo.EXPECT().SearchAccount(gomock.Any()).Return(
		[]businesslogic.Account{
			{ID: 301, Email: "testuser", PasswordHash: hash, PasswordSalt: salt, AccountStatusID: businesslogic.AccountStatusActivated},
		}, nil,
	)

	err := businesslogic.AuthenticateUser(username, password, userRepo)

	assert.Nil(t, err, "should not result in error if password hash matches")
}

func TestAuthenticateSuspendedUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	username := "testuser"
	password := "testpassword"
	userRepo := mock_businesslogic.NewMockIAccountRepository(mockCtrl)

	salt := util.GenerateSalt([]byte(password))
	hash := util.GenerateHash(salt, []byte(password))

	// specifies behavior
	userRepo.EXPECT().SearchAccount(gomock.Any()).Return(
		[]businesslogic.Account{
			{ID: 301, Email: "testuser", PasswordHash: hash, PasswordSalt: salt, AccountStatusID: businesslogic.AccountStatusSuspended},
		}, nil,
	)

	err := businesslogic.AuthenticateUser(username, password, userRepo)

	assert.NotNil(t, err, "should result in error if account is suspended")
}
