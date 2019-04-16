package organizer

import (
	"encoding/json"
	"fmt"
	"github.com/DancesportSoftware/das/auth"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
	"strconv"
)

// OrganizerLeadTagServer is a virtual server that handles request of searching and assigning lead tag
type OrganizerLeadTagServer struct {
	auth                        auth.IAuthenticationStrategy
	compRepo                    businesslogic.ICompetitionRepository
	partnershipCompEntryService businesslogic.PartnershipCompetitionEntryService
}

func NewOrganizerLeadTagServer(authenticaion auth.IAuthenticationStrategy, compRepo businesslogic.ICompetitionRepository, partnershipCompEntryService businesslogic.PartnershipCompetitionEntryService) OrganizerLeadTagServer {
	return OrganizerLeadTagServer{
		auth:                        authenticaion,
		compRepo:                    compRepo,
		partnershipCompEntryService: partnershipCompEntryService,
	}
}

// GetAllLeadEntries returns all the leads at the specified competition
func (server OrganizerLeadTagServer) GetAllLeadEntries(w http.ResponseWriter, r *http.Request) {
	formErr := r.ParseForm()
	if formErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, "bad form data", formErr)
		return
	}

	compId, convErr := strconv.Atoi(r.FormValue("competitionId"))
	if convErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, "invalid competition ID", nil)
		return
	}

	competitions, searchCompErr := server.compRepo.SearchCompetition(businesslogic.SearchCompetitionCriteria{ID: compId})
	if searchCompErr != nil || len(competitions) != 1 {
		util.RespondJsonResult(w, http.StatusBadRequest, fmt.Sprintf("cannot find competition with ID = %v", compId), nil)
		return
	}

	entries, err := server.partnershipCompEntryService.GetAllLeadEntries(competitions[0])
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	output := make([]viewmodel.AthleteCompetitionEntryViewModel, 0)
	for _, each := range entries {
		output = append(output, viewmodel.AthleteCompetitionEntryToViewModel(each))
	}

	jsonResp, _ := json.Marshal(output)
	w.Write(jsonResp)
}
