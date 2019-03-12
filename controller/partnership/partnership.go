package partnership

import (
	"encoding/json"
	"fmt"
	"github.com/DancesportSoftware/das/auth"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
	"time"
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
	PartnershipID int  `json:"partnershipId"`
	Favorite      bool `json:"favorite"`
}

// PUT /api/v1.0/athlete/partnership
func (server PartnershipServer) UpdatePartnershipHandler(w http.ResponseWriter, r *http.Request) {
	currentUser, _ := server.GetCurrentUser(r)
	if currentUser.ID == 0 {
		util.RespondJsonResult(w, http.StatusUnauthorized, "not authorized", nil)
		return
	}

	updateDTO := new(updatePartnership)
	if parseErr := util.ParseRequestBodyData(r, updateDTO); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, parseErr.Error())
		return
	}

	partnerships, searchErr := server.IPartnershipRepository.SearchPartnership(businesslogic.SearchPartnershipCriteria{PartnershipID: updateDTO.PartnershipID})
	if searchErr != nil || len(partnerships) != 1 {
		util.RespondJsonResult(w, http.StatusNotFound, fmt.Sprintf("cannot find partnership with ID = %v", updateDTO.PartnershipID), nil)
		return
	}

	// TODO [technical debt]: this should be in businesslogic
	//		check if current user owns the partnership, if not, then it's unauthorized. if yes, update the corresponding favorite
	if partnerships[0].HasAthlete(currentUser.ID) {
		if partnerships[0].Lead.ID == currentUser.ID {
			partnerships[0].FavoriteByLead = updateDTO.Favorite
		} else {
			partnerships[0].FavoriteByFollow = updateDTO.Favorite
		}
		partnerships[0].DateTimeUpdated = time.Now()
		if updateErr := server.IPartnershipRepository.UpdatePartnership(partnerships[0]); updateErr != nil {
			util.RespondJsonResult(w, http.StatusInternalServerError, updateErr.Error(), nil)
			return
		} else {
			util.RespondJsonResult(w, http.StatusOK, "partnership is updated successfully", nil)
			return
		}
	} else {
		util.RespondJsonResult(w, http.StatusUnauthorized, "not authorized to make changes to this partnership", nil)
		return
	}
}
