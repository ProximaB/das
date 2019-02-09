package partnership

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/auth"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
)

// PartnershipServer serves requests that are related with partnership
type PartnershipServer struct {
	auth.IAuthenticationStrategy
	businesslogic.IAccountRepository
	businesslogic.IPartnershipRepository
}

// SearchPartnershipHandler handles the request
//	GET /api/partnership
func (server PartnershipServer) SearchPartnershipHandler(w http.ResponseWriter, r *http.Request) {
	currentUser, _ := server.GetCurrentUser(r)
	if currentUser.ID == 0 || !currentUser.HasRole(businesslogic.AccountTypeAthlete) {
		util.RespondJsonResult(w, http.StatusUnauthorized, "not authorized", nil)
		return
	}

	partnerships, err := server.SearchPartnership(
		businesslogic.SearchPartnershipCriteria{AccountID: currentUser.ID})
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP500ErrorRetrievingData, err.Error())
		return
	}

	data := make([]viewmodel.Partnership, 0)
	for _, each := range partnerships {
		data = append(data, viewmodel.PartnershipDataModelToViewModel(currentUser, each))
	}
	output, _ := json.Marshal(data)
	w.Write(output)

}

type updatePartnership struct {
	PartnershipID int  `json:"partnership"`
	Favorite      bool `json:"favorite"`
}

// PUT /api/partnership
func (server PartnershipServer) UpdatePartnershipHandler(w http.ResponseWriter, r *http.Request) {
	account, _ := server.GetCurrentUser(r)
	if account.ID == 0 {
		util.RespondJsonResult(w, http.StatusUnauthorized, "not authorized", nil)
		return
	}

	updateDTO := new(updatePartnership)
	if parseErr := util.ParseRequestBodyData(r, updateDTO); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, parseErr.Error())
		return
	}
}
