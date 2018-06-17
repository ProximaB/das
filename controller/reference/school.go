package reference

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic/reference"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
)

type SchoolServer struct {
	referencebll.ISchoolRepository
}

// GET /api/reference/school
func (server SchoolServer) SearchSchoolHandler(w http.ResponseWriter, r *http.Request) {
	criteria := new(referencebll.SearchSchoolCriteria)

	if parseErr := util.ParseRequestData(r, criteria); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP_400_INVALID_REQUEST_DATA, parseErr.Error())
		return
	}

	if schools, err := server.SearchSchool(*criteria); err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP_500_ERROR_RETRIEVING_DATA, err.Error())
		return
	} else {
		data := make([]viewmodel.School, 0)
		for _, each := range schools {
			data = append(data, viewmodel.SchoolDataModelToViewModel(each))
		}

		output, _ := json.Marshal(data)
		w.Write(output)
	}
}

// POST /api/reference/school
func (server SchoolServer) CreateSchoolHandler(w http.ResponseWriter, r *http.Request) {}

// PUT /api/reference/school
func (server SchoolServer) UpdateSchoolHandler(w http.ResponseWriter, r *http.Request) {}

// DELETE /api/reference/school
func (server SchoolServer) DeleteSchoolHandler(w http.ResponseWriter, r *http.Request) {}
