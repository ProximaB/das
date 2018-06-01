package reference

import (
	"github.com/yubing24/das/businesslogic/reference"
	"github.com/yubing24/das/controller/util"
	"github.com/yubing24/das/viewmodel"
	"encoding/json"
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

type CountryServer struct {
	reference.ICountryRepository
}

// POST /api/reference/country
func (server CountryServer) CreateCountryHandler(w http.ResponseWriter, r *http.Request) {
	payload := new(viewmodel.CreateCountry)
	var err error

	if err = util.ParseRequestData(r, payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	country := payload.ToDataModel()

	if err := server.ICountryRepository.CreateCountry(&country); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	util.RespondJsonResult(w, http.StatusOK, "success", nil)
}

// DELETE /api/reference/country
func (server CountryServer) DeleteCountryHandler(w http.ResponseWriter, r *http.Request) {
	deleteDTO := new(viewmodel.DeleteCountry)
	err := util.ParseRequestData(r, deleteDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, viewmodel.RESTAPIResult{Message: err.Error()})
		return
	}

	country := reference.Country{
		ID: deleteDTO.CountryID,
	}

	err = server.ICountryRepository.DeleteCountry(country)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, viewmodel.RESTAPIResult{Message: err.Error()})
		return
	}
	fmt.Fprintln(w, viewmodel.RESTAPIResult{Message: "success"})
	return
}

// PUT /api/reference/country
func (server CountryServer) UpdateCountryHandler(w http.ResponseWriter, r *http.Request) {
	updateDTO := new(viewmodel.UpdateCountry)
	err := util.ParseRequestData(r, updateDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, viewmodel.RESTAPIResult{Message: err.Error()})
		return
	}

	country := reference.Country{
		ID:           updateDTO.CountryID,
		Name:         updateDTO.Name,
		Abbreviation: updateDTO.Abbreviation,
	}

	err = server.ICountryRepository.UpdateCountry(country)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, viewmodel.RESTAPIResult{Message: err.Error()})
		return
	}
	fmt.Fprintln(w, viewmodel.RESTAPIResult{Message: "success"})
	return
}

// GET /api/reference/country
func (server CountryServer) SearchCountryHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	searchDTO := new(reference.SearchCountryCriteria)
	if err := util.ParseRequestData(r, searchDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	countries, err := server.ICountryRepository.SearchCountry(searchDTO)
	if err != nil {
		log.Errorf(ctx, "error in searching Country: %v", err)
		util.RespondJsonResult(w, http.StatusInternalServerError, "cannot get countries", nil)
		return
	}

	output, _ := json.Marshal(viewmodel.CountriesToViewModel(countries))
	w.Write(output)
}
