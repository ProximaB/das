package reference

import (
	"github.com/DancesportSoftware/das/businesslogic/reference"
	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

var divisionRepo = PostgresDivisionRepository{

	Database:   nil,
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

func TestPostgresDivisionRepository_SearchDivision(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	divisionRepo.Database = db
	mock.ExpectQuery("SELECT")
	divisions, _ := divisionRepo.SearchDivision(&reference.SearchDivisionCriteria{})

	assert.Zero(t, len(divisions))
}
