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
	"time"
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
	AccountType int       `json:"accounttype"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	FirstName   string    `json:"firstname"`
	MiddleNames string    `json:"middlenames"`
	LastName    string    `json:"lastname"`
	DateOfBirth time.Time `json:"dateofbirth"`
	Gender      int       `json:"gender"`
	Password    string    `json:"password"`
	ToSAccepted bool      `json:"tosaccepted"`
	PPAccepted  bool      `json:"ppaccepted"`
	ByGuardian  bool      `json:"byguardian"`
	Signature   string    `json:"signature"`
}

type Login struct {
	Email    string `json:"username"`
	Password string `json:"password"`
}
