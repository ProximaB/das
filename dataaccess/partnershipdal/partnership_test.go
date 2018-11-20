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

package partnershipdal_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/partnershipdal"
	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var repo = partnershipdal.PostgresPartnershipRepository{
	Database:   nil,
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

func TestPostgresPartnershipRepository_CreatePartnership(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	repo.Database = db

	partnership := businesslogic.Partnership{
		LeadID:           12,
		FollowID:         14,
		FavoriteByFollow: true,
		FavoriteByLead:   true,
		DateTimeCreated:  time.Now(),
		DateTimeUpdated:  time.Now(),
	}

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO DAS.PARTNERSHIP (LEAD_ID, FOLLOW_ID, SAME_SEX_ID,
			STATUS_ID, FAVORITE_BY_LEAD, FAVORITE_BY_FOLLOW, COMPETITIONS_ATTENDED, EVENTS_ATTENDED, CREATE_USER_ID, `)
	mock.ExpectCommit()
	err := repo.CreatePartnership(&partnership)

	assert.Nil(t, err, "should at least return an empty array")
}

func TestPostgresPartnershipRepository_SearchPartnership(t *testing.T) {

}

func TestPostgresPartnershipRepository_UpdatePartnership(t *testing.T) {

}
