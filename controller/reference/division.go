package reference

import (
	"encoding/json"
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/controller/util"
	"github.com/ProximaB/das/viewmodel"
	"net/http"
)

type DivisionServer struct {
	businesslogic.IDivisionRepository
}

func (server DivisionServer) SearchDivisionHandler(w http.ResponseWriter, r *http.Request) {
	criteria := new(businesslogic.SearchDivisionCriteria)
	if parseErr := util.ParseRequestData(r, criteria); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, "invalid request data", parseErr.Error())
		return
	}

	if divisions, err := server.IDivisionRepository.SearchDivision(*criteria); err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP500ErrorRetrievingData, err.Error())
		return
	} else {
		data := make([]viewmodel.DivisionViewModel, 0)
		for _, each := range divisions {
			view := viewmodel.DivisionViewModel{
				ID:         each.ID,
				Name:       each.Name,
				Federation: each.FederationID,
			}
			data = append(data, view)
		}
		output, _ := json.Marshal(data)
		w.Write(output)
	}

}

func (server DivisionServer) CreateDivisionHandler(w http.ResponseWriter, r *http.Request) {}
func (server DivisionServer) UpdateDivisionHandler(w http.ResponseWriter, r *http.Request) {}
func (server DivisionServer) DeleteDivisionHandler(w http.ResponseWriter, r *http.Request) {}
