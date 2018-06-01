package partnership

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/controller/util"
	"github.com/yubing24/das/viewmodel"
	"encoding/json"
	"net/http"
)

type PartnershipServer struct {
	businesslogic.IAccountRepository
	businesslogic.IPartnershipRepository
}

// GET /api/partnership
func (server PartnershipServer) SearchPartnershipHandler(w http.ResponseWriter, r *http.Request) {
	account, _ := util.GetCurrentUser(r, server.IAccountRepository)
	if account.ID == 0 || account.AccountTypeID != businesslogic.ACCOUNT_TYPE_ATHLETE {
		util.RespondJsonResult(w, http.StatusUnauthorized, "not authorized", nil)
		return
	}

	partnerships, err := server.SearchPartnership(
		&businesslogic.SearchPartnershipCriteria{LeadID: account.ID, FollowID: account.ID})
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP_500_ERROR_RETRIEVING_DATA, err.Error())
		return
	}

	data := make([]viewmodel.Partnership, 0)
	for _, each := range partnerships {
		data = append(data, viewmodel.PartnershipDataModelToViewModel(each))
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
	account, _ := util.GetCurrentUser(r, server.IAccountRepository)
	if account.ID == 0 {
		util.RespondJsonResult(w, http.StatusUnauthorized, "not authorized", nil)
		return
	}

	updateDTO := new(updatePartnership)
	if parseErr := util.ParseRequestBodyData(r, updateDTO); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP_400_INVALID_REQUEST_DATA, parseErr.Error())
		return
	}
}
