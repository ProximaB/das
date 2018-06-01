package account

import (
	"github.com/yubing24/das/businesslogic/reference"
	"github.com/yubing24/das/viewmodel"
	"encoding/json"
	"net/http"
)

// this structure will allow controller-based caching in the future, if needed
type GenderServer struct {
	reference.IGenderRepository
}

// GET /api/account/gender
func (handler GenderServer) GetAccountGenderHandler(w http.ResponseWriter, r *http.Request) {
	data := make([]viewmodel.Gender, 0)
	genders, _ := handler.IGenderRepository.GetAllGenders()
	for _, each := range genders {
		data = append(data, viewmodel.GenderDataModelToViewModel(each))
	}
	output, _ := json.Marshal(data)
	w.Write(output)
}
