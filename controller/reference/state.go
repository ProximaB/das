package reference

import (
	"encoding/json"
	"fmt"
	"github.com/DancesportSoftware/das/businesslogic/reference"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
)

type StateServer struct {
	reference.IStateRepository
}

// GET /api/reference/state
func (server StateServer) SearchStateHandler(w http.ResponseWriter, r *http.Request) {
	criteria := new(reference.SearchStateCriteria)
	err := util.ParseRequestData(r, criteria)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, viewmodel.RESTAPIResult{Message: err.Error()})
		return
	}
	states, err := server.IStateRepository.SearchState(*criteria)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, viewmodel.RESTAPIResult{Message: err.Error()})
		return
	}
	output := make([]viewmodel.State, 0)
	for _, each := range states {
		output = append(output, viewmodel.StateDataModelToViewModel(each))
	}
	data, _ := json.Marshal(output)
	w.Write(data)
}

// POST /api/reference/state
func (server StateServer) CreateStateHandler(w http.ResponseWriter, r *http.Request) {
	util.RespondJsonResult(w, http.StatusNotImplemented, "not implemented", nil)
}

// PUT /api/reference/state
func (server StateServer) UpdateStateHandler(w http.ResponseWriter, r *http.Request) {
	util.RespondJsonResult(w, http.StatusNotImplemented, "not implemented", nil)
}

// DELETE /api/reference/state
func (server StateServer) DeleteStateHandler(w http.ResponseWriter, r *http.Request) {
	util.RespondJsonResult(w, http.StatusNotImplemented, "not implemented", nil)
}
