package referencedal_test

import (
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/dataaccess/referencedal"
	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

var stateRepository = referencedal.PostgresStateRepository{
	Database:   nil,
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

func TestPostgresStateRepository_SearchState(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	stateRepository.Database = db

	rows := sqlmock.NewRows(
		[]string{
			"ID", "NAME", "ADDRESS", "CITY_ID", "WEBSITE", "CREATE_USER_ID", "DATETIME_CREATED", "UPDATE_USER_ID", "DATETIME_UPDATED",
		},
	).AddRow(1, "Kanopy", "WI", 8, "www.example.com", 3, time.Now(), 4, time.Now())

	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	states, _ := stateRepository.SearchState(businesslogic.SearchStateCriteria{})

	assert.NotZero(t, len(states))
}

func TestPostgresStateRepository_CreateState(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	stateRepository.Database = db

	args := businesslogic.State{Name: "Commonwealth", CountryID: 32}

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO DAS.STATE (NAME, ABBREVIATION, COUNTRY_ID, CREATE_USER_ID, DATETIME_CREATED,
		UPDATE_USER_ID, DATETIME_UPDATED)`)
	mock.ExpectCommit()

	err = stateRepository.CreateState(&args)
	assert.Nil(t, err, "should insert new state without error")
}

func TestPostgresStateRepository_DeleteState(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	stateRepository.Database = db

	args := businesslogic.State{ID: 22, Name: "Commonwealth", CountryID: 32}

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM DAS.STATE`).WillReturnResult(sqlmock.NewResult(22, 1))
	mock.ExpectCommit()

	err = stateRepository.DeleteState(args)
	assert.Nil(t, err, "should delete new state without error")
}

func TestPostgresStateRepository_UpdateState(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	stateRepository.Database = db

	args := businesslogic.State{ID: 18, Name: "Commonwealth", CountryID: 32}

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE DAS.STATE`).WillReturnResult(sqlmock.NewResult(18, 1))
	mock.ExpectCommit()

	err = stateRepository.UpdateState(args)
	assert.Nil(t, err, "should update new state without error")
}
