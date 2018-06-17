package reference

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic/reference"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
)

type DanceServer struct {
	referencebll.IDanceRepository
}

// GET /api/reference/dance
func (server DanceServer) SearchDanceHandler(w http.ResponseWriter, r *http.Request) {
	criteria := new(referencebll.SearchDanceCriteria)
	if parseErr := util.ParseRequestData(r, criteria); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, "invalid request data", parseErr.Error())
		return
	}

	if dances, err := server.SearchDance(*criteria); err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, "error in retrieving dances", err.Error())
		return
	} else {
		data := make([]viewmodel.Dance, 0)
		for _, each := range dances {
			view := viewmodel.Dance{
				ID:           each.ID,
				Name:         each.Name,
				StyleID:      each.StyleID,
				Abbreviation: each.Abbreviation,
			}
			data = append(data, view)
		}
		output, _ := json.Marshal(data)
		w.Write(output)
	}
}
func (server DanceServer) CreateDanceHandler(w http.ResponseWriter, r *http.Request) {}
func (server DanceServer) UpdateDanceHandler(w http.ResponseWriter, r *http.Request) {}
func (server DanceServer) DeleteDanceHandler(w http.ResponseWriter, r *http.Request) {}
