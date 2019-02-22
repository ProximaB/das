package controller

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
)

type EntryServer struct {
	businesslogic.IEventRepository
	businesslogic.IPartnershipEventEntryRepository
}

// GetCompetitionEntryHandler handles the request
//	GET /api/competition/entries
// Public view for competitive event entry
func (server EntryServer) GetCompetitionEntryHandler(w http.ResponseWriter, r *http.Request) {
	criteria := new(viewmodel.SearchEntryForm)
	if parseErr := util.ParseRequestData(r, criteria); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, parseErr.Error())
		return
	}

	util.RespondJsonResult(w, http.StatusNotImplemented, "not implemented", nil)
	return
}
