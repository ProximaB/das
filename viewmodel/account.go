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
	"errors"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/businesslogic/reference"
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

type CreateAccount struct {
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Password    string `json:"password"`
	ToSAccepted bool   `json:"tosaccepted"`
	PPAccepted  bool   `json:"ppaccepted"`
}

func (dto CreateAccount) Validate() error {
	if len(dto.FirstName) < 2 || len(dto.LastName) < 2 {
		return errors.New("name is too short")
	}
	if len(dto.FirstName) > 18 || len(dto.LastName) > 18 {
		return errors.New("name is too long")
	}
	if len(dto.Email) < 5 {
		return errors.New("invalid email address")
	}
	if len(dto.Phone) < 3 {
		return errors.New("invalid phone number")
	}
	return nil
}

func (dto CreateAccount) ToAccountModel() businesslogic.Account {
	account := businesslogic.Account{
		FirstName:             dto.FirstName,
		LastName:              dto.LastName,
		UserGenderID:          reference.GENDER_UNKNOWN,
		Email:                 dto.Email,
		Phone:                 dto.Phone,
		ToSAccepted:           true,
		PrivacyPolicyAccepted: true,
	}
	return account
}

type Login struct {
	Email    string `json:"username"`
	Password string `json:"password"`
}
