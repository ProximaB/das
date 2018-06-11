package referencedal_test

import (
	"testing"
	"time"

	"github.com/DancesportSoftware/das/businesslogic/reference"

	"github.com/DancesportSoftware/das/dataaccess/referencedal"
	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var ageRepository = referencedal.PostgresAgeRepository{
	Database:   nil,
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var age = referencebll.Age{
	Name:            "Adult",
	AgeMaximum:      99,
	AgeMinimum:      19,
	Enforced:        false,
	DivisionID:      7,
	CreateUserID:    nil,
	DateTimeCreated: time.Now(),
	UpdateUserID:    nil,
	DateTimeUpdated: time.Now(),
}

func TestPostgresAgeRepository_SearchAge(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows(
		[]string{
			"ID", "NAME", "DESCRIPTION", "DIVISION_ID", "ENFORCED", "MINIMUM_AGE", "MAXIMUM_AGE",
			"CREATE_USER_ID", "DATETIME_CREATED", "UPDATE_USER_ID", "DATETIME_UPDATED"}).
		AddRow(
			1, "Pre Teen I", "Pre Teen of USA Dance", 1, false, 4, 6,
			2, time.Now(), 3, time.Now())
	mock.ExpectQuery(`SELECT ID, 	NAME, DESCRIPTION, DIVISION_ID, ENFORCED, MINIMUM_AGE, MAXIMUM_AGE, 
			CREATE_USER_ID,	DATETIME_CREATED, UPDATE_USER_ID, DATETIME_UPDATED FROM DAS.AGE`).WillReturnRows(rows)

	result, err := ageRepository.SearchAge(referencebll.SearchAgeCriteria{})
	if result != nil || err == nil {
		t.Errorf("should halt when search criteria or data source is nil")
	}

	ageRepository.Database = db

	ageRepository.SearchAge(referencebll.SearchAgeCriteria{
		DivisionID: 1,
		AgeID:      3,
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s\n", err)
	}
}

func TestPostgresAgeRepository_CreateAge(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	ageRepository.Database = db

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO DAS.AGE (NAME, DESCRIPTION, DIVISION_ID, ENFORCED, MINIMUM_AGE, MAXIMUM_AGE, 
			CREATE_USER_ID, DATETIME_CREATED, UPDATE_USER_ID, DATETIME_UPDATED)`)
	mock.ExpectCommit()
	err = ageRepository.CreateAge(&age)

	assert.Nil(t, err, "should insert a new age category without error")
}

func TestPostgresAgeRepository_DeleteAge(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ageRepository.Database = db
	args := referencebll.Age{ID: 1, Name: "Adult"}

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM DAS.AGE`).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = ageRepository.DeleteAge(args)
	assert.Nil(t, err, "should delete age without error")
}

func TestPostgresAgeRepository_UpdateAge(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ageRepository.Database = db
	args := referencebll.Age{ID: 1, Name: "Adult"}

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE DAS.AGE`).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = ageRepository.UpdateAge(args)
	assert.Nil(t, err, "should update age without error")
}
