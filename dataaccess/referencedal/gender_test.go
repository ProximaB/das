package referencedal_test

import (
	"github.com/ProximaB/das/dataaccess/common"
	"github.com/ProximaB/das/dataaccess/referencedal"
	"github.com/Masterminds/squirrel"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

var genderRepository = referencedal.PostgresGenderRepository{
	Database:   nil,
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

func TestPostgresGenderRepository_GetAllGenders(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows(
		[]string{
			common.ColumnPrimaryKey,
			common.COL_NAME,
			common.ColumnAbbreviation,
			common.COL_DESCRIPTION,
			common.ColumnDateTimeCreated,
			common.ColumnDateTimeUpdated,
		},
	).AddRow(
		1, "Female", "F", "Biologicially female", time.Now(), time.Now(),
	).AddRow(
		2, "Male", "M", "Biologically male", time.Now(), time.Now(),
	)

	mock.ExpectQuery(`SELECT ID, NAME, ABBREVIATION, DESCRIPTION, DATETIME_CREATED, DATETIME_UPDATED
FROM DAS.GENDER`).WillReturnRows(rows)
	genderRepository.Database = db
	genderRepository.GetAllGenders()

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s\n", err)
	}
}
