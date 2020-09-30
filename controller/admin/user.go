package admin

import (
	"encoding/json"
	"github.com/ProximaB/das/auth"
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/controller/util"
	"github.com/ProximaB/das/viewmodel"
	"net/http"
)

type AdminUserManagementServer struct {
	auth.IAuthenticationStrategy
	accountRepo businesslogic.IAccountRepository
}

func NewAdminUserManagementServer(auth auth.IAuthenticationStrategy, accountRepo businesslogic.IAccountRepository) AdminUserManagementServer {
	return AdminUserManagementServer{
		auth,
		accountRepo,
	}
}

// GET /api/v1/admin/user
func (server AdminUserManagementServer) SearchUserHandler(w http.ResponseWriter, r *http.Request) {
	currentUser, _ := server.GetCurrentUser(r)
	if !currentUser.HasRole(businesslogic.AccountTypeAdministrator) {
		util.RespondJsonResult(w, http.StatusBadRequest, "Not authorized to search user accounts", nil)
		return
	}
	searchCriteriaDTO := new(viewmodel.SearchAccountDTO)
	parseErr := util.ParseRequestData(r, searchCriteriaDTO)

	if parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, "Invalid search criteria data", nil)
		return
	}

	criteria := businesslogic.SearchAccountCriteria{}
	searchCriteriaDTO.Populate(&criteria)

	results, err := server.accountRepo.SearchAccount(criteria)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, "An internal ", nil)
		return
	}

	data := make([]viewmodel.AccountDTO, 0)
	for _, each := range results {
		dto := viewmodel.AccountDTO{}
		dto.Extract(each)
		data = append(data, dto)
	}

	output, _ := json.Marshal(data)
	w.Write(output)
}
