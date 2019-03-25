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

type SearchAccountDTO struct {
	FirstName string `schema:"firstName"`
	LastName  string `schema:"lastName"`
	RoleID    int    `schema:"roleId"`
	Email     string `schema:"email"`
	Phone     string `schema:"phone"`
}

func (dto SearchAccountDTO) Populate(criteria *businesslogic.SearchAccountCriteria) {
	criteria.AccountType = dto.RoleID
	criteria.FirstName = dto.FirstName
	criteria.LastName = dto.LastName
	criteria.Email = dto.Email
	criteria.Phone = dto.Phone
}

type AccountDTO struct {
	FirstName       string    `json:"firstName"`
	LastName        string    `json:"lastName"`
	Email           string    `json:"email"`
	Phone           string    `json:"phone"`
	Roles           []int     `json:"roles"`
	DateTimeCreated time.Time `json:"createdOn"`
	DateTimeUpdated time.Time `json:"updatedOn"`
}

func (dto *AccountDTO) Extract(account businesslogic.Account) {
	dto.FirstName = account.FirstName
	dto.LastName = account.LastName
	dto.Email = account.Email
	dto.Phone = account.Phone
	dto.Roles = account.GetRoles()
	dto.DateTimeCreated = account.DateTimeCreated
	dto.DateTimeUpdated = account.DateTimeModified
}

// CreateAccountDTO is the JSON payload for request POST /api/v1.0/account/register
type CreateAccountDTO struct {
	Email       string `json:"email" validate:"regexp=^[0-9a-z]+@[0-9a-z]+(\\.[0-9a-z]+)+$"`
	Phone       string `json:"phone"`
	FirstName   string `json:"firstname" validate:"nonzero"`
	LastName    string `json:"lastname" validate:"nonzero"`
	ToSAccepted bool   `json:"tosaccepted" validate:"true"`
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

// AthleteTinyViewModel is the minimum data of an athlete
type AthleteTinyViewModel struct {
	UID       string `json:"uid"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func AthleteToTinyViewModel(athlete businesslogic.Account) AthleteTinyViewModel {
	return AthleteTinyViewModel{
		UID:       athlete.UID,
		FirstName: athlete.FirstName,
		LastName:  athlete.LastName,
	}
}
