package reference

import (
	"github.com/yubing24/das/businesslogic/reference"
	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

var cityRepository = PostgresCityRepository{
	Database:   nil,
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var city = reference.City{
	Name:            "Test City",
	StateID:         1,
	DateTimeCreated: time.Now(),
	DateTimeUpdated: time.Now(),
}

func TestPostgresCityRepository_SearchCity(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	cityRepository.Database = db
	mock.ExpectQuery(`SELECT ID, NAME, STATE_ID, CREATE_USER_ID, DATETIME_CREATED, UPDATE_USER_ID, DATETIME_UPDATED FROM DAS.CITY`)
	cities, err := cityRepository.SearchCity(&reference.SearchCityCriteria{})

	assert.NotZero(t, len(cities), "should retrieve cities that were populated to Database")
	assert.Nil(t, err, "schema for DAS.CITY should be up to date")
}

func TestPostgresCityRepository_CreateCity(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	cityRepository.Database = db
	mock.ExpectExec("DELETE")
	cityRepository.DeleteCity(city)
	err = cityRepository.CreateCity(&city)
	assert.Nil(t, err, "should be able to create a new city")
	assert.NotZero(t, city.CityID, "should retrieve the ID after creation")
	cityRepository.DeleteCity(city)

}

func TestPostgresCityRepository_UpdateCity(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	cityRepository.Database = db
	mock.ExpectExec("UPDATE")

	cityRepository.UpdateCity(city)

}

func TestPostgresCityRepository_DeleteCity(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	cityRepository.Database = db
	mock.ExpectQuery("SELECT")
	cities, _ := cityRepository.SearchCity(&reference.SearchCityCriteria{Name: "Test City"})
	assert.EqualValues(t, 1, len(cities), "data should exist in Database before deletion")
	err = cityRepository.DeleteCity(city)
	assert.Nil(t, err, "should be able to delete the city that was created")
}
