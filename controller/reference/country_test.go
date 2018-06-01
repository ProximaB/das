package reference

import (
	"github.com/yubing24/das/businesslogic/reference"
	"github.com/yubing24/das/mock/businesslogic/reference"
	"github.com/yubing24/das/viewmodel"
	"encoding/json"
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
	server := CountryServer{mockedCountryRepo}

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

	assert.Nil(t, err, "should return a list of countries")
	w.Flush()

	// test with bad param
	// TODO: this is not working for some reason. Probably URL encoding
	/*
		query := r.URL.Query()
		query.Add("badparam", "indeed")
		r.RequestURI = query.Encode()
		r.Form.Add("badparam","verybad")
		log.Println(r.Form)
		log.Println(r.URL.String())
		server.SearchCountryHandler(w, r)
		log.Println(w.Body.String())
		assert.EqualValues(t, http.StatusBadRequest, w.Code, "should receive HTTP 400 when sending a bad request")
	*/
}
