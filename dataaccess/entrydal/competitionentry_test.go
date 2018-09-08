// Dancesport Application System (DAS)
// Copyright (C) 2018 Yubing Hou
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package entrydal

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/util"
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
		AthleteID: 12,
		CompetitionEntry: businesslogic.BaseCompetitionEntry{
			CompetitionID:    12,
			CheckInIndicator: true,
			DateTimeCheckIn:  time.Now(),
		},
	}

	err := athleteCompEntryRepo.CreateEntry(&entry)
	assert.NotNil(t, err, dalutil.ErrorNilDatabase)

	athleteCompEntryRepo.Database = db

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO DAS.COMPETITION_ENTRY_ATHLETE (COMPETITION_ID, ATHLETE_ID, CHECKIN_IND,
		CHECKIN_DATETIME, PAYMENT_IND, CREATE_USER_ID, DATETIME_CREATED, UPDATE_USER_ID, DATETIME_UPDATED)`)
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
		CompetitionEntry: businesslogic.BaseCompetitionEntry{
			CompetitionID:    4,
			CheckInIndicator: false,
			DateTimeCheckIn:  time.Now(),
		},
		PartnershipID: 12,
	}

	err := partnershipCompEntryRepo.CreateEntry(&entry)
	assert.NotNil(t, err, dalutil.ErrorNilDatabase)

	partnershipCompEntryRepo.Database = db

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO DAS.COMPETITION_ENTRY_PARTNERSHIP (COMPETITION_ID, PARTNERSHIP_ID, CHECKIN_IND,
		CHECKIN_DATETIME, CREATE_USER_ID, DATETIME_CREATED, UPDATE_USER_ID, DATETIME_UPDATED)`)
	mock.ExpectCommit()
	err = partnershipCompEntryRepo.CreateEntry(&entry)

	assert.Nil(t, err, "should insert legitimate PartnershipCompetitionEntry data without error")
}
