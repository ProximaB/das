package reference

import (
	"encoding/json"
	"fmt"
	"github.com/DancesportSoftware/das/businesslogic/reference"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"log"
	"net/http"
)

type CountryServer struct {
	referencebll.ICountryRepository
}

// POST /api/referencedal/country
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

// DELETE /api/referencedal/country
func (server CountryServer) DeleteCountryHandler(w http.ResponseWriter, r *http.Request) {
	deleteDTO := new(viewmodel.DeleteCountry)
	err := util.ParseRequestData(r, deleteDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, viewmodel.RESTAPIResult{Message: err.Error()})
		return
	}

	country := referencebll.Country{
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

// PUT /api/referencedal/country
func (server CountryServer) UpdateCountryHandler(w http.ResponseWriter, r *http.Request) {
	updateDTO := new(viewmodel.UpdateCountry)
	err := util.ParseRequestData(r, updateDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, viewmodel.RESTAPIResult{Message: err.Error()})
		return
	}

	country := referencebll.Country{
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

// GET /api/referencedal/country
func (server CountryServer) SearchCountryHandler(w http.ResponseWriter, r *http.Request) {
	if server.ICountryRepository == nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, "data source for SearchCountryHandler is not specified", nil)
		return
	}

	searchDTO := new(referencebll.SearchCountryCriteria)
	if err := util.ParseRequestData(r, searchDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	countries, err := server.ICountryRepository.SearchCountry(*searchDTO)
	if err != nil {
		log.Printf("error in searching Country: %v", err)
		util.RespondJsonResult(w, http.StatusInternalServerError, "cannot get countries", nil)
		return
	}

	output, _ := json.Marshal(viewmodel.CountriesToViewModel(countries))
	w.Write(output)
}
