package reference

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
)

type SchoolServer struct {
	businesslogic.ISchoolRepository
}

// GET /api/reference/school
func (server SchoolServer) SearchSchoolHandler(w http.ResponseWriter, r *http.Request) {
	criteria := new(businesslogic.SearchSchoolCriteria)

	if parseErr := util.ParseRequestData(r, criteria); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, parseErr.Error())
		return
	}

	if schools, err := server.SearchSchool(*criteria); err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP500ErrorRetrievingData, err.Error())
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
