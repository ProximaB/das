package account

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic/reference"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
)

// this structure will allow controller-based caching in the future, if needed
type GenderServer struct {
	referencebll.IGenderRepository
}

// GetAccountGenderHandler handles request
//		GET /api/account/gender
// No parameter is required for this request.
//
// Sample returned result:
//	[
// 		{"id":1,"name":"Female"},
// 		{"id":2,"name":"Male"}
// 	]
func (handler GenderServer) GetAccountGenderHandler(w http.ResponseWriter, r *http.Request) {
	data := make([]viewmodel.Gender, 0)
	genders, _ := handler.IGenderRepository.GetAllGenders()
	for _, each := range genders {
		data = append(data, viewmodel.GenderDataModelToViewModel(each))
	}
	output, _ := json.Marshal(data)
	w.Write(output)
}
