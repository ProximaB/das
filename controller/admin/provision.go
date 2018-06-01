package admin

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/controller/util"
	"github.com/yubing24/das/viewmodel"
	"net/http"
)

type OrganizerProvisionServer struct {
	businesslogic.IAccountRepository
	businesslogic.IOrganizerProvisionRepository
}

// PUT /api/admin/organizer/organizer
func (server OrganizerProvisionServer) UpdateOrganizerProvisionHandler(w http.ResponseWriter, r *http.Request) {
	updateDTO := new(viewmodel.UpdateProvision)
	parseErr := util.ParseRequestBodyData(r, updateDTO)
	if parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, "invalid request data", nil)
		return
	}

	organizer := businesslogic.GetAccountByUUID(updateDTO.OrganizerID, server.IAccountRepository)
	provisions, _ := server.SearchOrganizerProvision(&businesslogic.SearchOrganizerProvisionCriteria{OrganizerID: organizer.ID})

	record := provisions[0]
	// TODO: finish implementing the data update

	err := server.UpdateOrganizerProvision(record)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	util.RespondJsonResult(w, http.StatusOK, "success", nil)
}
