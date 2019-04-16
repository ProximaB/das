package account

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
)

// AccountTypeServer is a micro-server that serves requests that ask for available
// account types in DAS
type AccountTypeServer struct {
	businesslogic.IAccountTypeRepository
}

// GetAccountTypeHandler handles the request
//	GET /api/account/type
// No parameter is required for this request.
//
// Sample returned result:
//	[
// 		{"id":1,"name":"Athlete"},
// 		{"id":2,"name":"Adjudicator"},
// 		{"id":3,"name":"Scrutineer"},
// 		{"id":4,"name":"Organizer"},
// 		{"id":5,"name":"Deck Captain"},
// 		{"id":6,"name":"Emcee"}
// 	]
func (server AccountTypeServer) GetAccountTypeHandler(w http.ResponseWriter, r *http.Request) {
	data := make([]viewmodel.AccountTypePublicView, 0)
	types, _ := server.GetAccountTypes()
	for _, each := range types {
		data = append(data, viewmodel.NewAccountTypePublicView(each))
	}
	output, _ := json.Marshal(data)
	w.Write(output)
}
