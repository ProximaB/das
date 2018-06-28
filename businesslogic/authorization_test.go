// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

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
