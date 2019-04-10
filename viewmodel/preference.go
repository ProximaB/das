package viewmodel

import "github.com/DancesportSoftware/das/businesslogic"

type UserPreferenceViewModel struct {
	DefaultRole int `json:"defaultRole"`
}

func UserPreferenceDataModelToViewModel(model businesslogic.UserPreference) UserPreferenceViewModel {
	return UserPreferenceViewModel{
		DefaultRole: model.DefaultRole,
	}
}
