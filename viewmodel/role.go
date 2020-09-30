package viewmodel

import "github.com/ProximaB/das/businesslogic"

type AccountRoleDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func AccountRoleToAccountRoleDTO(role businesslogic.AccountRole) AccountRoleDTO {
	dto := AccountRoleDTO{
		ID: role.AccountTypeID,
	}

	switch role.AccountTypeID {
	case businesslogic.AccountTypeAthlete:
		dto.Name = "Athlete"
		break
	case businesslogic.AccountTypeAdjudicator:
		dto.Name = "Adjudicator"
		break
	case businesslogic.AccountTypeScrutineer:
		dto.Name = "Scrutineer"
		break
	case businesslogic.AccountTypeOrganizer:
		dto.Name = "Organizer"
		break
	case businesslogic.AccountTypeDeckCaptain:
		dto.Name = "Deck Captain"
		break
	case businesslogic.AccountTypeEmcee:
		dto.Name = "Emcee"
		break
	case businesslogic.AccountTypeAdministrator:
		dto.Name = "Administrator"
		break

	}
	return dto
}
