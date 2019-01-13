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

package viewmodel

import (
	"github.com/DancesportSoftware/das/businesslogic"
)

type AccountType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func AccountTypeDataModelToViewModel(dm businesslogic.AccountType) AccountType {
	return AccountType{
		ID:   dm.ID,
		Name: dm.Name,
	}
}

// CreateAccountDTO is the JSON payload for request POST /api/v1.0/account/register
type CreateAccountDTO struct {
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	ToSAccepted bool   `json:"tosaccepted"`
	PPAccepted  bool   `json:"ppaccepted"`
}

func (dto CreateAccountDTO) ToAccountModel() businesslogic.Account {
	account := businesslogic.Account{
		FirstName:             dto.FirstName,
		LastName:              dto.LastName,
		UserGenderID:          businesslogic.GENDER_UNKNOWN,
		Email:                 dto.Email,
		Phone:                 dto.Phone,
		ToSAccepted:           true,
		PrivacyPolicyAccepted: true,
	}
	return account
}
