package controller

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/dataaccess"
	"net/http"
)

var competitionEntryRepository = dataaccess.PostgresCompetitionEntryRepository{}
var eventEntryRepository = dataaccess.PostgresEventEntryRepository{}

// GET /api/entries
// Public view for competitive event entry
func getCompetitiveBallroomEventEntryHandler(w http.ResponseWriter, r *http.Request) {
	criteria := new(businesslogic.SearchEventEntryCriteria)
	if parseErr := util.ParseRequestData(r, criteria); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP_400_INVALID_REQUEST_DATA, parseErr.Error())
		return
	}

	entries, err := eventEntryRepository.SearchEventEntry(criteria)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP_500_ERROR_RETRIEVING_DATA, err.Error())
		return
	}

	output, _ := json.Marshal(entries)
	w.Write(output)

}
