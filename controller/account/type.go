package account

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/viewmodel"
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
