package reference

import (
	"github.com/yubing24/das/businesslogic/reference"
	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

var danceRepo = PostgresDanceRepository{
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
	mock.ExpectQuery("SELECT")
	dances, _ := danceRepo.SearchDance(&reference.SearchDanceCriteria{})

	assert.Zero(t, len(dances))
}
