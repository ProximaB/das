package database

import (
	"github.com/DancesportSoftware/das/dataaccess/accountdal"
	"github.com/DancesportSoftware/das/dataaccess/competition"
	"github.com/DancesportSoftware/das/dataaccess/entrydal"
	"github.com/DancesportSoftware/das/dataaccess/eventdal"
	"github.com/DancesportSoftware/das/dataaccess/organizer"
	"github.com/DancesportSoftware/das/dataaccess/partnershipdal"
	"github.com/DancesportSoftware/das/dataaccess/provision"
	"github.com/DancesportSoftware/das/dataaccess/referencedal"
	"github.com/Masterminds/squirrel"
)

//======= Reference data repositories

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

//======== end of reference data repositories

//======== begin of account repositories
var AccountRepository = accountdal.PostgresAccountRepository{
	SQLBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var AccountRoleRepository = accountdal.PostgresAccountRoleRepository{
	SQLBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var UserPreferenceRepository = accountdal.PostgresUserPreferenceRepository{
	SQLBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var AccountTypeRepository = accountdal.PostgresAccountTypeRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var RoleApplicationRepository = accountdal.PostgresRoleApplicationRepository{
	SQLBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

//========= end of account repositories

var PartnershipRepository = partnershipdal.PostgresPartnershipRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var PartnershipRoleRepository = partnershipdal.PostgresPartnershipRoleRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var PartnershipRequestRepository = partnershipdal.PostgresPartnershipRequestRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var PartnershipRequestStatusRepository = partnershipdal.PostgresPartnershipRequestStatusRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var PartnershipRequestBlacklistRepository = partnershipdal.PostgresPartnershipRequestBlacklistRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var PartnershipRequestBlacklistReasonRepository = partnershipdal.PostgresPartnershipRequestBlacklistReasonRepository{
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

var CompetitionOfficialRepository = organizer.PostgresCompetitionOfficialRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var CompetitionOfficialInvitationRepository = organizer.PostgresCompetitionOfficialInvitationRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var AthleteCompetitionEntryRepository = entrydal.PostgresAthleteCompetitionEntryRepository{
	SQLBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var PartnershipCompetitionEntryRepository = entrydal.PostgresPartnershipCompetitionEntryRepository{
	SQLBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var EventRepository = eventdal.PostgresEventRepository{
	SQLBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var EventMetaRepository = eventdal.PostgresEventMetaRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var EventDanceRepository = eventdal.PostgresEventDanceRepository{
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var CompetitionEventTemplateRepository = eventdal.PostgresCompetitionEventTemplateRepository{
	SQLBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var AthleteEventEntryRepository = entrydal.PostgresAthleteEventEntryRepository{
	SQLBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var PartnershipEventEntryRepository = entrydal.PostgresPartnershipEventEntryRepository{
	SQLBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}
