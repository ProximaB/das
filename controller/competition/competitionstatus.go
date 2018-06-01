package competition

import (
	"github.com/yubing24/das/controller/util"
	"github.com/yubing24/das/dataaccess"
	"github.com/yubing24/das/viewmodel"
	"encoding/json"
	"net/http"
)

var competitionStatusRepository = dataaccess.PostgresCompetitionStatusRepository{}

// GET /api/competition/status
func getCompetitionStatusHandler(w http.ResponseWriter, r *http.Request) {
	status, err := competitionStatusRepository.GetCompetitionStatus()
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP_500_ERROR_RETRIEVING_DATA, err.Error())
		return
	} else {
		data := make([]viewmodel.CompetitionStatus, 0)
		for _, each := range status {
			data = append(data, viewmodel.CompetitionStatusDataModelToViewModel(each))
		}
		output, _ := json.Marshal(data)
		w.Write(output)
	}
}
