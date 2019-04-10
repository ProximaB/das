package referencedal_test

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/referencedal"
	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

var danceRepo = referencedal.PostgresDanceRepository{
	Database:   nil,
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

func TestPostgresDanceRepository_SearchDance(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	danceRepo.Database = db

	rows := sqlmock.NewRows(
		[]string{
			"ID", "NAME", "ABBREVIATION", "DESCRIPTION", "STYLE_ID", "CREATE_USER_ID", "DATETIME_CREATED",
			"UPDATE_USER_ID", "DATETIME_UPDATED"}).
		AddRow(
			1, "Waltz", "W", "International Waltz", 1, 2, time.Now(), 3, time.Now()).
		AddRow(
			2, "Tango", "T", "International Tango", 1, 2, time.Now(), 3, time.Now())

	mock.ExpectQuery(`SELECT ID, NAME, ABBREVIATION, DESCRIPTION, STYLE_ID, CREATE_USER_ID, 
			DATETIME_CREATED, UPDATE_USER_ID, DATETIME_UPDATED FROM DAS.DANCE`).WillReturnRows(rows)
	dances, _ := danceRepo.SearchDance(businesslogic.SearchDanceCriteria{})

	assert.EqualValues(t, 2, len(dances), "search with empty criteria should return all dances")
}

func TestPostgresDanceRepository_CreateDance(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	danceRepo.Database = db

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO DAS.DANCE`)
	mock.ExpectCommit()

	args := businesslogic.Dance{ID: 3, Name: "Foxtrot", StyleID: 4, DateTimeUpdated: time.Now()}

	err = danceRepo.CreateDance(&args)

	assert.Nil(t, err, "should create dance without error")
}

func TestPostgresDanceRepository_DeleteDance(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	danceRepo.Database = db

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM DAS.DANCE`).WillReturnResult(sqlmock.NewResult(3, 1))
	mock.ExpectCommit()

	args := businesslogic.Dance{ID: 3, Name: "Foxtrot", StyleID: 4, DateTimeUpdated: time.Now()}

	err = danceRepo.DeleteDance(args)

	assert.Nil(t, err, "should delete dance without error")
}

func TestPostgresDanceRepository_UpdateDance(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	danceRepo.Database = db

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE DAS.DANCE SET`).WillReturnResult(sqlmock.NewResult(3, 1))
	mock.ExpectCommit()

	args := businesslogic.Dance{ID: 3, Name: "Foxtrot", StyleID: 4, DateTimeUpdated: time.Now()}

	err = danceRepo.UpdateDance(args)

	assert.Nil(t, err, "should update dance without error")

}
