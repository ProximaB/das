package reference

import (
	"encoding/json"
	"fmt"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
)

type FederationServer struct {
	businesslogic.IFederationRepository
}

// GET /api/reference/federation
func (server FederationServer) SearchFederationHandler(w http.ResponseWriter, r *http.Request) {
	criteria := new(businesslogic.SearchFederationCriteria)
	if err := util.ParseRequestData(r, criteria); err != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, err.Error())
		return
	}

	federations, err := server.IFederationRepository.SearchFederation(*criteria)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, viewmodel.RESTAPIResult{Message: err.Error()})
		return
	}

	dtos := make([]viewmodel.Federation, 0)
	for _, each := range federations {
		dtos = append(dtos, viewmodel.Federation{
			ID:           each.ID,
			Name:         each.Name,
			Abbreviation: each.Abbreviation,
		})
	}
	output, _ := json.Marshal(dtos)
	w.Write(output)
}

// POST /api/reference/federation
func (server FederationServer) CreateFederationHandler(w http.ResponseWriter, r *http.Request) {}

// DELETE /api/reference/federation
func (server FederationServer) DeleteFederationHandler(w http.ResponseWriter, r *http.Request) {}

// PUT /api/reference/federation
func (server FederationServer) UpdateFederationHandler(w http.ResponseWriter, r *http.Request) {}
