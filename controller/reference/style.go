package reference

import (
	"encoding/json"
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/controller/util"
	"github.com/ProximaB/das/viewmodel"
	"net/http"
)

type StyleServer struct {
	businesslogic.IStyleRepository
}

// GET /api/reference/style
func (server StyleServer) SearchStyleHandler(w http.ResponseWriter, r *http.Request) {
	criteria := new(businesslogic.SearchStyleCriteria)
	if parseErr := util.ParseRequestData(r, criteria); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, "invalid request data", parseErr.Error())
		return
	}

	if styles, err := server.IStyleRepository.SearchStyle(*criteria); err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, "error in retrieving styles", err.Error())
		return
	} else {
		data := make([]viewmodel.Style, 0)
		for _, each := range styles {
			viewmodel := viewmodel.Style{
				ID:   each.ID,
				Name: each.Name,
			}
			data = append(data, viewmodel)
		}

		output, _ := json.Marshal(data)
		w.Write(output)
	}

}

func (server StyleServer) CreateStyleHandler(w http.ResponseWriter, r *http.Request) {}
func (server StyleServer) UpdateStyleHandler(w http.ResponseWriter, r *http.Request) {}
func (server StyleServer) DeleteStyleHandler(w http.ResponseWriter, r *http.Request) {}
