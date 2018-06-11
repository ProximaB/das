package viewmodel

import "github.com/DancesportSoftware/das/businesslogic/reference"

type Studio struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	CityID  int    `json:"city"`
	Website string `json:"website"`
}

func StudioDataModelToViewModel(dm referencebll.Studio) Studio {
	return Studio{
		ID:      dm.ID,
		Name:    dm.Name,
		Address: dm.Address,
		CityID:  dm.CityID,
		Website: dm.Website,
	}
}
