package viewmodel

import "github.com/DancesportSoftware/das/businesslogic"

type AccountRoleDTO struct {
	ID int `json:"roleId"`
}

func AccountRoleToAccountRoleDTO(role businesslogic.AccountRole) AccountRoleDTO {
	return AccountRoleDTO{
		ID: role.AccountTypeID,
	}
}
