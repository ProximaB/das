package reference_test

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic/reference"
	refcontroller "github.com/DancesportSoftware/das/controller/reference"
	"github.com/DancesportSoftware/das/mock/businesslogic/reference"
	"github.com/DancesportSoftware/das/viewmodel"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchCountryHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockedCountryRepo := mock_reference.NewMockICountryRepository(mockCtrl)
	server := refcontroller.CountryServer{ICountryRepository: mockedCountryRepo}

	r, _ := http.NewRequest(http.MethodGet, "/api/reference/country", nil)
	w := httptest.NewRecorder()

	mockedCountryRepo.EXPECT().SearchCountry(gomock.Any()).Return([]reference.Country{
		{ID: 1, Name: "United States"},
		{ID: 2, Name: "United Kingdom"},
	}, nil)

	// test with zero param
	server.SearchCountryHandler(w, r)
	countries := make([]viewmodel.Country, 0)
	err := json.Unmarshal([]byte(w.Body.String()), &countries)

	assert.EqualValues(t, 2, len(countries), "search country without parameter should get all countries")
	assert.EqualValues(t, 1, countries[0].ID, "ID should match")
	assert.Nil(t, err, "should return a list of countries")
	w.Flush()

	// test with bad param
	// TODO: this is not working for some reason. Probably URL encoding
	r, _ = http.NewRequest(http.MethodGet, "/api/reference/country", nil)
	query := r.URL.Query()
	query.Add("badparam", "indeed")
	r.RequestURI = query.Encode()
	//r.Form.Add("badparam","verybad")

	mockedCountryRepo.EXPECT().SearchCountry(gomock.Any()).Return([]reference.Country{
		{ID: 1, Name: "United States"},
		{ID: 2, Name: "United Kingdom"},
	}, nil)
	w = httptest.NewRecorder()

	server.SearchCountryHandler(w, r)

	res := viewmodel.RESTAPIResult{}

	err = json.Unmarshal([]byte(w.Body.String()), &res)
	assert.EqualValues(t, http.StatusBadRequest, res.Status, "should receive HTTP 400 when sending a bad request")
	w.Flush()
}

func TestCountryServer_CreateCountryHandler(t *testing.T) {
}

func TestCountryServer_DeleteCountryHandler(t *testing.T) {

}

func TestCountryServer_SearchCountryHandler(t *testing.T) {
	server := refcontroller.CountryServer{}
	w := httptest.NewRecorder()
	var req *http.Request
	var res viewmodel.RESTAPIResult
	var err error

	// null repository
	req, _ = http.NewRequest(http.MethodGet, "/api/reference/country", nil)
	server.SearchCountryHandler(w, req)
	res = viewmodel.RESTAPIResult{}
	json.Unmarshal([]byte(w.Body.String()), &res)
	assert.EqualValues(t, http.StatusInternalServerError, res.Status, "should return internal server error when repository is null")
	w.Flush()

	// initialize the controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockedCountryRepo := mock_reference.NewMockICountryRepository(mockCtrl)
	server.ICountryRepository = mockedCountryRepo
	mockedCountryRepo.EXPECT().SearchCountry(gomock.Any()).Return([]reference.Country{
		{ID: 1, Name: "United States"},
		{ID: 2, Name: "United Kingdom"},
	}, nil)

	// empty parameter
	server.SearchCountryHandler(w, req)
	countries := make([]viewmodel.Country, 0)
	err = json.Unmarshal([]byte(w.Body.String()), countries)
	assert.Nil(t, err, "should not throw an error with correct setup")
	assert.EqualValues(t, 2, len(countries), "")

}
