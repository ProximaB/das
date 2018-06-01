package viewmodel

import "github.com/yubing24/das/businesslogic/reference"

type School struct {
	SchoolID int    `json:"id"`
	Name     string `json:"name"`
	CityID   int    `json:"city"`
}

func SchoolDataModelToViewModel(school reference.School) School {
	return School{
		SchoolID: school.ID,
		Name:     school.Name,
		CityID:   school.CityID,
	}
}
