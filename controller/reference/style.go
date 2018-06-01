package reference

import (
	"github.com/yubing24/das/businesslogic/reference"
	"github.com/yubing24/das/controller/util"
	"github.com/yubing24/das/viewmodel"
	"encoding/json"
	"net/http"
)

type StyleServer struct {
	reference.IStyleRepository
}

// GET /api/reference/style
func (server StyleServer) SearchStyleHandler(w http.ResponseWriter, r *http.Request) {
	criteria := new(reference.SearchStyleCriteria)
	if parseErr := util.ParseRequestData(r, criteria); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, "invalid request data", parseErr.Error())
		return
	}

	if styles, err := server.IStyleRepository.SearchStyle(criteria); err != nil {
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
