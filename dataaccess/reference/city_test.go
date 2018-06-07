package reference

import (
	"github.com/DancesportSoftware/das/businesslogic/reference"
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

	rows := sqlmock.NewRows(
		[]string{"ID", "NAME", "STATE_ID", "CREATE_USER_ID", "DATETIME_CREATED", "UPDATE_USER_ID", "DATETIME_UPDATED"},
	).AddRow(
		1, "Madison", 2, 3, time.Now(), 3, time.Now(),
	).AddRow(
		2, "Milwaukee", 4, 5, time.Now(), 4, time.Now(),
	)

	mock.ExpectQuery(`SELECT ID, NAME, STATE_ID, CREATE_USER_ID, DATETIME_CREATED, UPDATE_USER_ID, DATETIME_UPDATED FROM DAS.CITY`).WillReturnRows(rows)
	cities, err := cityRepository.SearchCity(reference.SearchCityCriteria{})

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

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO DAS.CITY (NAME, STATE_ID, CREATE_USER_ID, DATETIME_CREATED, UPDATE_USER_ID, DATETIME_UPDATED) `)
	mock.ExpectCommit()
	err = cityRepository.CreateCity(&city)

	assert.Nil(t, err, "should be able to create a new city")
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
