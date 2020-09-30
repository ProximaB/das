package partnershipdal_test

import (
	"github.com/ProximaB/das/dataaccess/partnershipdal"
	"github.com/Masterminds/squirrel"
	"testing"
)

var repo = partnershipdal.PostgresPartnershipRepository{
	Database:   nil,
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

// TODO: test is rigged
func TestPostgresPartnershipRepository_CreatePartnership(t *testing.T) {
	/*db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	repo.Database = db

	partnership := businesslogic.Partnership{
		Lead:             businesslogic.Account{ID: 12},
		Follow:           businesslogic.Account{ID: 14},
		FavoriteByFollow: true,
		FavoriteByLead:   true,
		DateTimeCreated:  time.Now(),
		DateTimeUpdated:  time.Now(),
	}

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO DAS.PARTNERSHIP (LEAD_ID, FOLLOW_ID, SAME_SEX_ID,
			STATUS_ID, FAVORITE_BY_LEAD, FAVORITE_BY_FOLLOW, COMPETITIONS_ATTENDED, EVENTS_ATTENDED, CREATE_USER_ID, `)
	mock.ExpectCommit()
	err = repo.CreatePartnership(&partnership)

	assert.Nil(t, err, "should at least return an empty array")*/
}
