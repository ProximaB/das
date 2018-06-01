package reference

import (
	"github.com/yubing24/das/businesslogic/reference"
	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

var schoolRepository = PostgresSchoolRepository{

	Database:   nil,
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

func TestDasSchoolRepository_SearchSchool(t *testing.T) {
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
	schools, _ := schoolRepository.SearchSchool(&reference.SearchSchoolCriteria{})

	assert.NotZero(t, len(schools))
}
