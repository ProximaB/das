package viewmodel

import (
	"github.com/DancesportSoftware/das/businesslogic/reference"
)

type Gender struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GenderDataModelToViewModel(gender referencebll.Gender) Gender {
	return Gender{
		ID:   gender.ID,
		Name: gender.Name,
	}
}
