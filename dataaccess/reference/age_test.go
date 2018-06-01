package reference

import (
	"testing"
	"time"

	"github.com/yubing24/das/businesslogic/reference"

	"github.com/Masterminds/squirrel"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var ageRepository = PostgresAgeRepository{
	Database:   nil,
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

func TestPostgresAgeRepository_SearchAge(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows(
		[]string{
			"id",
			"name",
			"description",
			"division_id",
			"enforced", "minimum_age",
			"MAXIMUM_AGE",
			"create_user_id",
			"datetime_created",
			"update_user_id",
			"datetime_updated"}).
		AddRow(
			1,
			"Pre Teen I",
			"Pre Teen of USA Dance",
			1,
			false,
			4,
			6,
			2,
			time.Now(),
			3,
			time.Now())
	mock.ExpectQuery(`SELECT 
			ID, 
			NAME, 
			DESCRIPTION, 
			DIVISION_ID, 
			ENFORCED,
			MINIMUM_AGE, 
			MAXIMUM_AGE, 
			CREATE_USER_ID, 
			DATETIME_CREATED, 
			UPDATE_USER_ID, 
			DATETIME_UPDATED 
			FROM DAS.AGE`).WillReturnRows(rows)

	result, err := ageRepository.SearchAge(nil)
	if result != nil || err == nil {
		t.Errorf("should halt when search criteria or data source is nil")
	}

	ageRepository.Database = db

	ageRepository.SearchAge(&reference.SearchAgeCriteria{
		DivisionID: 1,
		AgeID:      3,
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s\n", err)
	}
}
