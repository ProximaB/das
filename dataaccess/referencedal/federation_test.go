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

var federationRepo = referencedal.PostgresFederationRepository{
	Database:   nil,
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

func TestPostgresFederationRepository_SearchFederation(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	federationRepo.Database = db
	rows := sqlmock.NewRows(
		[]string{
			"ID",
			"NAME",
			"ABBREVIATION",
			"YEAR_FOUNDED",
			"COUNTRY_ID",
			"CREATE_USER_ID",
			"DATETIME_CREATED",
			"UPDATE_USER_ID",
			"DATETIME_UPDATED",
		},
	).AddRow(1,
		"Pre Teen I",
		"Pre",
		1233,
		8,
		2,
		time.Now(),
		3,
		time.Now())
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	federations, err := federationRepo.SearchFederation(businesslogic.SearchFederationCriteria{})
	assert.Nil(t, err, "should be able to read federation table")
	assert.NotZero(t, len(federations), "Database has more than 1 federation")
}

func TestPostgresFederationRepository_CreateFederation(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	federationRepo.Database = db

	args := businesslogic.Federation{Name: "SUPER DANCE", Description: "Super Dancesport Federation"}

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO DAS.FEDERATION (NAME, ABBREVIATION, DESCRIPTION, 
		YEAR_FOUNDED, COUNTRY_ID, CREATE_USER_ID, DATETIME_CREATED, UPDATE_USER_ID, DATETIME_UPDATED)`)
	mock.ExpectCommit()

	err = federationRepo.CreateFederation(&args)
	assert.Nil(t, err, "should insert new federation without error")
}

func TestPostgresFederationRepository_DeleteFederation(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	federationRepo.Database = db

	args := businesslogic.Federation{ID: 37, Name: "SUPER DANCE"}

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM DAS.FEDERATION`).WillReturnResult(sqlmock.NewResult(37, 1))
	mock.ExpectCommit()

	err = federationRepo.DeleteFederation(args)
	assert.Nil(t, err, "should delete federation without error")
}

func TestPostgresFederationRepository_UpdateFederation(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	federationRepo.Database = db

	args := businesslogic.Federation{ID: 37, Name: "SUPER DANCE"}

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE DAS.FEDERATION`).WillReturnResult(sqlmock.NewResult(37, 1))
	mock.ExpectCommit()

	err = federationRepo.UpdateFederation(args)
	assert.Nil(t, err, "should update federation without error")
}
