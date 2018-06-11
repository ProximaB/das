package viewmodel

import "github.com/DancesportSoftware/das/businesslogic/reference"

type State struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	CountryID    int    `json:"country"`
}

func StateDataModelToViewModel(dm referencebll.State) State {
	return State{
		ID:           dm.ID,
		Name:         dm.Name,
		Abbreviation: dm.Abbreviation,
		CountryID:    dm.CountryID,
	}
}
