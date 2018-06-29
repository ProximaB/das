// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package database

import (
	"github.com/DancesportSoftware/das/dataaccess/account"
	"github.com/DancesportSoftware/das/dataaccess/competition"
	"github.com/DancesportSoftware/das/dataaccess/entry"
	"github.com/DancesportSoftware/das/dataaccess/event"
	"github.com/DancesportSoftware/das/dataaccess/partnership"
	"github.com/DancesportSoftware/das/dataaccess/provision"
	"github.com/DancesportSoftware/das/dataaccess/referencedal"
	"github.com/Masterminds/squirrel"
)

var CountryRepository = referencedal.PostgresCountryRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var StateRepository = referencedal.PostgresStateRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var CityRepository = referencedal.PostgresCityRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var FederationRepository = referencedal.PostgresFederationRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var DivisionRepository = referencedal.PostgresDivisionRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var AgeRepository = referencedal.PostgresAgeRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var ProficiencyRepository = referencedal.PostgresProficiencyRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var StyleRepository = referencedal.PostgresStyleRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var DanceRepository = referencedal.PostgresDanceRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var SchoolRepository = referencedal.PostgresSchoolRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var StudioRepository = referencedal.PostgresStudioRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var AccountRepository = account.PostgresAccountRepository{
	SQLBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var AccountTypeRepository = account.PostgresAccountTypeRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var PartnershipRepository = partnership.PostgresPartnershipRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var PartnershipRequestRepository = partnership.PostgresPartnershipRequestRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var PartnershipRequestStatusRepository = partnership.PostgresPartnershipRequestStatusRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var PartnershipRequestBlacklistRepository = partnership.PostgresPartnershipRequestBlacklistRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var PartnershipRequestBlacklistReasonRepository = partnership.PostgresPartnershipRequestBlacklistReasonRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var GenderRepository = referencedal.PostgresGenderRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var OrganizerProvisionRepository = provision.PostgresOrganizerProvisionRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var OrganizerProvisionHistoryRepository = provision.PostgresOrganizerProvisionHistoryRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var CompetitionStatusRepository = competition.PostgresCompetitionStatusRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var CompetitionRepository = competition.PostgresCompetitionRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var AthleteCompetitionEntryRepository = entry.PostgresAthleteCompetitionEntryRepository{
	SQLBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var PartnershipCompetitionEntryRepository = entry.PostgresPartnershipCompetitionEntryRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var EventRepository = event.PostgresEventRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var EventMetaRepository = event.PostgresEventMetaRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var PartnershipEventEntryRepository = entry.PostgresPartnershipEventEntryRepository{
	SQLBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}
