package entrydal_test

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/entrydal"
	"github.com/DancesportSoftware/das/dataaccess/util"
	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

var partnershipEventEntryRepo = entrydal.PostgresPartnershipEventEntryRepository{
	Database:   nil,
	SQLBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

func TestPostgresPartnershipEventEntryRepository_CreatePartnershipEventEntry(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	checkin := time.Now()
	entry := businesslogic.PartnershipEventEntry{
		Couple:            businesslogic.Partnership{ID: 37},
		Event:             businesslogic.Event{ID: 991},
		DateTimeCheckedIn: &checkin,
	}

	err := partnershipEventEntryRepo.CreatePartnershipEventEntry(&entry)
	assert.NotNil(t, err, dalutil.ErrorNilDatabase)

	partnershipEventEntryRepo.Database = db

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO DAS.EVENT_ENTRY_PARTNERSHIP (EVENT_ID, PARTNERSHIP_ID, LEADTAG, CHECKIN_IND,
		CHECKIN_DATETIME, CREATE_USER_ID, DATETIME_CREATED, UPDATE_USER_ID, DATETIME_UPDATED)`)
	mock.ExpectCommit()

	err = partnershipEventEntryRepo.CreatePartnershipEventEntry(&entry)
	assert.Nil(t, err, "should insert legitimate PartnershipEventEntry data without error")
}
