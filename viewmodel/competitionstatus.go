package viewmodel

import "github.com/yubing24/das/businesslogic"

type CompetitionStatus struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func CompetitionStatusDataModelToViewModel(status businesslogic.CompetitionStatus) CompetitionStatus {
	return CompetitionStatus{
		ID:   status.ID,
		Name: status.Name,
	}
}
