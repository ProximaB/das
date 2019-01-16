package viewmodel

import "github.com/DancesportSoftware/das/businesslogic"

type CompetitionOfficialProfileDTO struct {
	Name          string `json:"name"`
	IsAdjudicator bool   `json:"isAdjudicator"`
	IsScrutineer  bool   `json:"isScrutineer"`
	IsDeckCaptain bool   `json:"isDeckCaptain"`
	IsEmcee       bool   `json:"isEmcee"`
}

func (dto *CompetitionOfficialProfileDTO) Populate(account businesslogic.Account) {
	dto.Name = account.FullName()
	dto.IsAdjudicator = account.HasRole(businesslogic.AccountTypeAdjudicator)
	dto.IsScrutineer = account.HasRole(businesslogic.AccountTypeScrutineer)
	dto.IsDeckCaptain = account.HasRole(businesslogic.AccountTypeDeckCaptain)
	dto.IsEmcee = account.HasRole(businesslogic.AccountTypeEmcee)
}
