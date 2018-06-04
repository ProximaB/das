package reference

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic/reference"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
)

type AgeServer struct {
	reference.IAgeRepository
}

func (server AgeServer) SearchAgeHandler(w http.ResponseWriter, r *http.Request) {
	criteria := new(reference.SearchAgeCriteria)
	if parseErr := util.ParseRequestData(r, criteria); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP_400_INVALID_REQUEST_DATA, parseErr.Error())
		return
	}

	if ages, err := server.SearchAge(*criteria); err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP_500_ERROR_RETRIEVING_DATA, err.Error())
		return
	} else {
		data := make([]viewmodel.Age, 0)
		for _, each := range ages {
			data = append(data, viewmodel.AgeDataModelToViewModel(each))
		}

		output, _ := json.Marshal(data)
		w.Write(output)
	}
}

func (server AgeServer) CreateAgeHandler(w http.ResponseWriter, r *http.Request) {

}

func (server AgeServer) UpdateAgeHandler(w http.ResponseWriter, r *http.Request) {

}

func (server AgeServer) DeleteAgeHandler(w http.ResponseWriter, r *http.Request) {

}
