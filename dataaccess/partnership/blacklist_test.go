package partnership

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

var blacklistRepository = PostgresPartnershipRequestBlacklistRepository{
	Database:   nil,
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var blacklist = businesslogic.PartnershipRequestBlacklistEntry{
	BlockedUserID: 1,
}

func TestPostgresPartnershipRequestBlacklistRepository_CreatePartnershipRequestBlacklist(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	blacklistRepository.Database = db

	mock.ExpectQuery(`SELECT ID, REPORTER_ID, BLOCKED_USER_ID, BLACKLIST_REASON_ID, DETAIL, 
		WHITELISTED_IND, CREATE_USER_ID, DATETIME_CREATED, UPDATE_USER_ID, DATETIME_UPDATED FROM 
		DAS.PARTNERSHIP_REQUEST_BLACKLIST`).WillReturnRows()
	results, err := blacklistRepository.SearchPartnershipRequestBlacklist(businesslogic.SearchPartnershipRequestBlacklistCriteria{})

	assert.Nil(t, err, "should not throw when empty parameter is provided")
	assert.NotNil(t, results, "should at least return an empty array")
}
