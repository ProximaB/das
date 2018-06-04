package competition

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
)

type CompetitionStatusServer struct {
	businesslogic.ICompetitionStatusRepository
}

// GET /api/competition/status
func (server CompetitionStatusServer) GetCompetitionStatusHandler(w http.ResponseWriter, r *http.Request) {
	status, err := server.GetCompetitionAllStatus()
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
