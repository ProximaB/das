package entrydal

// TODO: this cannot be tested because DATA-DOG's SQL mock does not support queryrow

/*
import (
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/dataaccess/util"
	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

var athleteCompEntryRepo = PostgresAthleteCompetitionEntryRepository{
	Database:   nil,
	SQLBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

func TestPostgresAthleteCompetitionEntryRepository_CreateAthleteCompetitionEntry(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	entry := businesslogic.AthleteCompetitionEntry{
		Athlete: businesslogic.Account{ID: 12},
		Competition: businesslogic.Competition{ID: 12},
		CheckedIn: true,
		DateTimeCheckedIn: time.Now(),
	}

	err := athleteCompEntryRepo.CreateEntry(&entry)
	assert.NotNil(t, err, dalutil.ErrorNilDatabase)

	athleteCompEntryRepo.Database = db

	mock.ExpectBegin()
	mock.ExpectQuery(`^INSERT INTO DAS.COMPETITION_ENTRY_ATHLETE (COMPETITION_ID,ATHLETE_ID,CHECKIN_IND,CHECKIN_DATETIME,PAYMENT_IND,CREATE_USER_ID,DATETIME_CREATED,UPDATE_USER_ID,DATETIME_UPDATED) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING ID`).
		WithArgs(entry.Competition.ID, entry.Athlete.ID, entry.CheckedIn, entry.DateTimeCheckedIn,entry.PaymentReceivedIndicator,entry.CreateUserID,entry.DateTimeCreated,entry.UpdateUserID,entry.DateTimeUpdated)
	mock.ExpectCommit()

	err = athleteCompEntryRepo.CreateEntry(&entry)

	assert.Nil(t, err, "should insert legitimate AthleteCompetitionEntry data without error")
}

var partnershipCompEntryRepo = PostgresPartnershipCompetitionEntryRepository{
	Database:   nil,
	SQLBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

func TestPostgresPartnershipCompetitionEntryRepository_CreatePartnershipCompetitionEntry(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	entry := businesslogic.PartnershipCompetitionEntry{
		Couple: businesslogic.Partnership{ID: 12},
		Competition: businesslogic.Competition{ID: 4},
		CheckedIn: false,
	}

	err := partnershipCompEntryRepo.CreateEntry(&entry)
	assert.NotNil(t, err, dalutil.ErrorNilDatabase)

	partnershipCompEntryRepo.Database = db

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO DAS.COMPETITION_ENTRY_PARTNERSHIP (COMPETITION_ID, PARTNERSHIP_ID, CHECKIN_IND,
		CHECKIN_DATETIME, CREATE_USER_ID, DATETIME_CREATED, UPDATE_USER_ID, DATETIME_UPDATED) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING ID`)
	mock.ExpectCommit()
	err = partnershipCompEntryRepo.CreateEntry(&entry)

	assert.Nil(t, err, "should insert legitimate PartnershipCompetitionEntry data without error")
}
*/
