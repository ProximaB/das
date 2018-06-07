package reference

import (
	"github.com/DancesportSoftware/das/businesslogic/reference"
	"github.com/Masterminds/squirrel"
	"github.com/go-errors/errors"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

var countryRepo = PostgresCountryRepository{
	Database:   nil,
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

func TestPostgresCountryRepository_CreateCountry(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	country := reference.Country{
		Name:         "United States",
		Abbreviation: "USA",
	}
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO DAS.COUNTRY").WithArgs(country.Name,
		country.Abbreviation, nil, country.DateTimeCreated, nil, country.DateTimeUpdated)
	mock.ExpectCommit()

	countryRepo.Database = db
	err = countryRepo.CreateCountry(&country)
	assert.Nil(t, err, "should create a new country")
}

func TestPostgresCountryRepository_SearchCountry(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows(
		[]string{
			"ID", "NAME", "ABBREVIATION", "CREATE_USER_ID", "DATETIME_CREATED", "UPDATE_USER_ID", "DATETIME_UPDATED",
		},
	).AddRow(
		1, "United States", "USA", 1, time.Now(), 1, time.Now(),
	).AddRow(
		2, "Canada", "CAN", 1, time.Now(), 1, time.Now(),
	)

	mock.ExpectQuery(`SELECT ID, NAME, ABBREVIATION, 
		CREATE_USER_ID,  DATETIME_CREATED, UPDATE_USER_ID,
		DATETIME_UPDATED FROM DAS.COUNTRY`).WillReturnRows(rows)

	countryRepo.Database = db
	countries, err := countryRepo.SearchCountry(reference.SearchCountryCriteria{})
	assert.Nil(t, err, "should get all countries")
	assert.EqualValues(t, len(countries), 2, "should return all countries")
}

func TestPostgresCountryRepository_DeleteCountry(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// TODO: yhou 2018-06-07: delete test with SQL Mock is not clear.....
	mock.ExpectBegin()
	mock.ExpectExec(`^DELETE FROM DAS.COUNTRY`).WillReturnResult(sqlmock.NewErrorResult(errors.New("")))
	mock.ExpectCommit()

	countryRepo.Database = db
	err = countryRepo.DeleteCountry(reference.Country{ID: 1})
	assert.Nil(t, err, "should delete country without error")
}
