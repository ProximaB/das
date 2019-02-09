package organizer

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/config/routes/middleware"
	"github.com/DancesportSoftware/das/controller/organizer"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiOrganizerCompetitionOfficialSearch = "/api/v1/organizer/competition/official/eligible"

var organizerCompetitionOfficialSearchServer = organizer.OrganizerCompetitionOfficialSearchServer{
	IAuthenticationStrategy: middleware.AuthenticationStrategy,
	IAccountRepository:      database.AccountRepository,
	IAccountRoleRepository:  database.AccountRoleRepository,
}

var SearchEligibleCompetitionOfficialController = util.DasController{
	Name:         "SearchEligibleCompetitionOfficialController",
	Description:  "Organzier search eligible officials for competition",
	Method:       http.MethodGet,
	Endpoint:     apiOrganizerCompetitionOfficialSearch,
	Handler:      organizerCompetitionOfficialSearchServer.SearchEligibleOfficialHandler,
	AllowedRoles: []int{businesslogic.AccountTypeOrganizer},
}
