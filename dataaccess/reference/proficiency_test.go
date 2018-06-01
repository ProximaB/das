package reference

import (
	"github.com/yubing24/das/businesslogic/reference"
	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

var proficiencyRepo = PostgresProficiencyRepository{

	Database:   nil,
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

func TestPostgresProficiencyRepository_SearchProficiency(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	proficiencyRepo.Database = db

	rows := sqlmock.NewRows(
		[]string{
			"ID", "NAME", "DIVISION_ID", "DESCRIPTION", "CREATE_USER_ID", "DATETIME_CREATED", "UPDATE_USER_ID", "DATETIME_UPDATED",
		},
	).AddRow(1, "Gold", 3, "USA DANCE Gold", 3, time.Now(), 4, time.Now())

	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	proficienies, _ := proficiencyRepo.SearchProficiency(&reference.SearchProficiencyCriteria{})

	assert.EqualValues(t, 1, len(proficienies))
}
