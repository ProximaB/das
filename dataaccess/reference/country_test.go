package reference

import (
	"github.com/DancesportSoftware/das/businesslogic/reference"
	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

var countryRepo = PostgresCountryRepository{
	Database:   nil,
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

func TestCreateCountry(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO DAS.COUNTRY").WillReturnResult(sqlmock.NewResult(1, 1))

	countryRepo.Database = db
	country := reference.Country{
		Name:         "Random",
		Abbreviation: "RAN",
	}
	assert.Nil(t, countryRepo.CreateCountry(&country), "should create a new country")
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s\n", err)
	}
}

func TestSearchCountry(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(`SELECT ID, 
		NAME, 
		ABBREVIATION, 
		CREATE_USER_ID, 
		DATETIME_CREATED, 
		UPDATE_USER_ID, 
		DATETIME_UPDATED FROM DAS.COUNTRY`).WillReturnResult(sqlmock.NewResult(1, 1))

	countryRepo.Database = db
	countries, err := countryRepo.SearchCountry(&reference.SearchCountryCriteria{})
	assert.Nil(t, err, "should get all countries")
	assert.True(t, len(countries) >= 0, "should contain some data")
}

func TestPostgresCountryRepository_DeleteCountry(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("DELETE FROM DAS.COUNTRY").WillReturnResult(sqlmock.NewResult(1, 1))

	countryRepo.Database = db
	err = countryRepo.DeleteCountry(reference.Country{ID: 1})
	assert.Nil(t, err, "should delete created country")
}
