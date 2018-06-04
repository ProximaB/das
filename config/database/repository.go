package database

import (
	"github.com/DancesportSoftware/das/dataaccess/account"
	"github.com/DancesportSoftware/das/dataaccess/competition"
	"github.com/DancesportSoftware/das/dataaccess/event"
	"github.com/DancesportSoftware/das/dataaccess/partnership"
	"github.com/DancesportSoftware/das/dataaccess/provision"
	"github.com/DancesportSoftware/das/dataaccess/reference"
	"github.com/Masterminds/squirrel"
)

var CountryRepository = reference.PostgresCountryRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var StateRepository = reference.PostgresStateRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var CityRepository = reference.PostgresCityRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var FederationRepository = reference.PostgresFederationRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var DivisionRepository = reference.PostgresDivisionRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var AgeRepository = reference.PostgresAgeRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var ProficiencyRepository = reference.PostgresProficiencyRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var StyleRepository = reference.PostgresStyleRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var DanceRepository = reference.PostgresDanceRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var SchoolRepository = reference.PostgresSchoolRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var StudioRepository = reference.PostgresStudioRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var AccountRepository = account.PostgresAccountRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
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

var GenderRepository = reference.PostgresGenderRepository{
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

var EventRepository = event.PostgresEventRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var EventMetaRepository = event.PostgresEventMetaRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}
