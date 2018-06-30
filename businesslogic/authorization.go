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

package businesslogic

import (
	"errors"
	"reflect"

	"github.com/DancesportSoftware/das/util"
)

func AuthenticateUser(email, password string, repo IAccountRepository) error {
	if len(email) < 1 || len(password) < 1 {
		return errors.New("username or password is missing")
	}

	accounts, err := repo.SearchAccount(SearchAccountCriteria{Email: email})
	if err != nil {
		return err
	}

	if len(accounts) != 1 {
		return errors.New("invalid credential")
	}

	if accounts[0].AccountStatusID == AccountStatusSuspended {
		return errors.New("account is suspended")
	}

	if accounts[0].AccountStatusID == AccountStatusLocked {
		return errors.New("account is locked")
	}

	expectedHash := util.GenerateHash(accounts[0].PasswordSalt, []byte(password))
	if reflect.DeepEqual(expectedHash, accounts[0].PasswordHash) {
		return nil // user is authenticated
	}
	return errors.New("incorrect username or password")
}
