package account

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/viewmodel"
	"encoding/json"
	"net/http"
)

type AccountTypeServer struct {
	businesslogic.IAccountTypeRepository
}

// GET /api/account/type
func (server AccountTypeServer) GetAccountTypeHandler(w http.ResponseWriter, r *http.Request) {
	data := make([]viewmodel.AccountType, 0)
	types, _ := server.GetAccountTypes()
	for _, each := range types {
		if each.ID != businesslogic.ACCOUNT_TYPE_ADMINISTRATOR {
			data = append(data, viewmodel.AccountTypeDataModelToViewModel(each))
		}
	}
	output, _ := json.Marshal(data)
	w.Write(output)
}
