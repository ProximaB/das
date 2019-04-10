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

var schoolRepository = referencedal.PostgresSchoolRepository{
	Database:   nil,
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

func TestPostgresSchoolRepository_SearchSchool(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	schoolRepository.Database = db

	rows := sqlmock.NewRows(
		[]string{
			"ID", "NAME", "CITY_ID", "CREATE_USER_ID", "DATETIME_CREATED", "UPDATE_USER_ID", "DATETIME_UPDATED",
		},
	).AddRow(1, "UW-Madison", 3, 3, time.Now(), 4, time.Now())

	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	schools, _ := schoolRepository.SearchSchool(businesslogic.SearchSchoolCriteria{})

	assert.NotZero(t, len(schools))
}

func TestPostgresSchoolRepository_CreateSchool(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	schoolRepository.Database = db

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO DAS.SCHOOL (NAME, CITY_ID, CREATE_USER_ID, DATETIME_CREATED,
		UPDATE_USER_ID, DATETIME_UPDATED)`)
	mock.ExpectCommit()
	args := businesslogic.School{Name: "Intergalactic College", CityID: 44}
	err = schoolRepository.CreateSchool(&args)

	assert.Nil(t, err, "should insert new school without error")
}

func TestPostgresSchoolRepository_DeleteSchool(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	schoolRepository.Database = db

	args := businesslogic.School{ID: 66, Name: "Intergalactic College", CityID: 44}
	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM DAS.SCHOOL`).WillReturnResult(sqlmock.NewResult(66, 1))
	mock.ExpectCommit()

	err = schoolRepository.DeleteSchool(args)
	assert.Nil(t, err, "should delete school without error")
}

func TestPostgresSchoolRepository_UpdateSchool(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	schoolRepository.Database = db

	args := businesslogic.School{ID: 66, Name: "Intergalactic College", CityID: 44}
	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE DAS.SCHOOL`).WillReturnResult(sqlmock.NewResult(66, 1))
	mock.ExpectCommit()

	err = schoolRepository.UpdateSchool(args)
	assert.Nil(t, err, "should update school without error")
}
