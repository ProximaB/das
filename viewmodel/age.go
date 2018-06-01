package viewmodel

import (
	"github.com/yubing24/das/businesslogic/reference"
)

type SearchAge struct {
	DivisionID int `schema:"division"`
	AgeID      int `schema:"id"`
}

type Age struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Division int    `json:"division"`
	Enforced bool   `json:"enforced"`
	Minimum  int    `json:"minimum"`
	Maximum  int    `json:"maximum"`
}

func AgeDataModelToViewModel(dm reference.Age) Age {
	return Age{
		ID:       dm.ID,
		Name:     dm.Name,
		Division: dm.DivisionID,
		Enforced: dm.Enforced,
		Minimum:  dm.AgeMinimum,
		Maximum:  dm.AgeMaximum,
	}
}
