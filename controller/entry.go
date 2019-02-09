package controller

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

type EntryServer struct {
	businesslogic.IEventRepository
	businesslogic.IPartnershipEventEntryRepository
}

// GET /api/entries
// Public view for competitive event entry
func (server EntryServer) getCompetitiveBallroomEventEntryHandler(w http.ResponseWriter, r *http.Request) {
	criteria := new(businesslogic.SearchPartnershipEventEntryCriteria)
	if parseErr := util.ParseRequestData(r, criteria); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, parseErr.Error())
		return
	}

	entries, err := server.SearchPartnershipEventEntry(*criteria)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP500ErrorRetrievingData, err.Error())
		return
	}

	output, _ := json.Marshal(entries)
	w.Write(output)

}
