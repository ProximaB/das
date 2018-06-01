package competition

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/controller/util"
	"github.com/yubing24/das/dataaccess"
	"github.com/yubing24/das/viewmodel"
	"encoding/json"
	"net/http"
	"strconv"
)

var competitionRepository = dataaccess.PostgresCompetitionRepository{}

// GET /api/public/competitions
// Search competition(s). This controller is invokable without authentication
func publicSearchCompetitionHandler(w http.ResponseWriter, r *http.Request) {
	searchDTO := new(businesslogic.SearchCompetitionCriteria)
	if parseErr := util.ParseRequestData(r, searchDTO); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP_400_INVALID_REQUEST_DATA, parseErr.Error())
		return
	} else {
		if competitions, err := competitionRepository.SearchCompetition(&businesslogic.SearchCompetitionCriteria{
			ID:       searchDTO.ID,
			Name:     searchDTO.Name,
			StatusID: searchDTO.StatusID,
		}); err != nil {
			util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP_500_ERROR_RETRIEVING_DATA, err.Error())
		} else {
			data := make([]viewmodel.Competition, 0)
			for _, each := range competitions {
				data = append(data, viewmodel.CompetitionDataModelToViewModel(each, businesslogic.ACCOUNT_TYPE_NOAUTH))
			}
			output, _ := json.Marshal(data)
			w.Write(output)
		}
	}
}

// GET /api/competition/federation
func getEventUniqueFederationsHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	compID, parseErr := strconv.Atoi(r.Form.Get("competition"))
	if parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP_400_INVALID_REQUEST_DATA, parseErr.Error())
		return
	}

	federations, err := businesslogic.GetEventUniqueFederations(compID)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP_500_ERROR_RETRIEVING_DATA, err.Error())
		return
	}

	data := make([]viewmodel.Federation, 0)
	for _, each := range federations {
		data = append(data, viewmodel.Federation{
			ID:           each.ID,
			Name:         each.Name,
			Abbreviation: each.Abbreviation,
		})
	}

	output, _ := json.Marshal(data)
	w.Write(output)
}

// GET /api/competition/division
func getEventUniqueDivisionsHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	compID, parseErr := strconv.Atoi(r.Form.Get("competition"))
	if parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP_400_INVALID_REQUEST_DATA, parseErr.Error())
		return
	}

	divisions, err := businesslogic.GetEventUniqueDivisions(compID)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP_500_ERROR_RETRIEVING_DATA, err.Error())
		return
	}

	data := make([]viewmodel.DivisionViewModel, 0)
	for _, each := range divisions {
		data = append(data, viewmodel.DivisionViewModel{
			ID:         each.ID,
			Name:       each.Name,
			Federation: each.FederationID,
		})
	}

	output, _ := json.Marshal(data)
	w.Write(output)
}

// GET /api/competition/age
func getEventUniqueAgesHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	compID, parseErr := strconv.Atoi(r.Form.Get("competition"))
	if parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP_400_INVALID_REQUEST_DATA, parseErr.Error())
		return
	}

	federations, err := businesslogic.GetEventUniqueAges(compID)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP_500_ERROR_RETRIEVING_DATA, err.Error())
		return
	}

	data := make([]viewmodel.Age, 0)
	for _, each := range federations {
		data = append(data, viewmodel.Age{
			ID:       each.ID,
			Name:     each.Name,
			Division: each.DivisionID,
			Maximum:  each.AgeMaximum,
			Minimum:  each.AgeMinimum,
		})
	}

	output, _ := json.Marshal(data)
	w.Write(output)
}

// GET /api/competition/proficiency
func getEventUniqueProficienciesHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	compID, parseErr := strconv.Atoi(r.Form.Get("competition"))
	if parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP_400_INVALID_REQUEST_DATA, parseErr.Error())
		return
	}

	federations, err := businesslogic.GetEventUniqueProficiencies(compID)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP_500_ERROR_RETRIEVING_DATA, err.Error())
		return
	}

	data := make([]viewmodel.Proficiency, 0)
	for _, each := range federations {
		data = append(data, viewmodel.ProficiencyDataModelToViewModel(each))
	}

	output, _ := json.Marshal(data)
	w.Write(output)
}

// GET /api/competition/style
func getEventUniqueStylesHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	compID, parseErr := strconv.Atoi(r.Form.Get("competition"))
	if parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP_400_INVALID_REQUEST_DATA, parseErr.Error())
		return
	}

	federations, err := businesslogic.GetEventUniqueStyles(compID)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP_500_ERROR_RETRIEVING_DATA, err.Error())
		return
	}

	data := make([]viewmodel.Style, 0)
	for _, each := range federations {
		data = append(data, viewmodel.Style{
			ID:   each.ID,
			Name: each.Name,
		})
	}

	output, _ := json.Marshal(data)
	w.Write(output)
}
