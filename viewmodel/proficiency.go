package viewmodel

import "github.com/yubing24/das/businesslogic/reference"

type Proficiency struct {
	ProficiencyID int    `json:"id"`
	Name          string `json:"name"`
	Division      int    `json:"division"`
}

func ProficiencyDataModelToViewModel(dm reference.Proficiency) Proficiency {
	return Proficiency{
		ProficiencyID: dm.ID,
		Name:          dm.Name,
		Division:      dm.DivisionID,
	}
}
