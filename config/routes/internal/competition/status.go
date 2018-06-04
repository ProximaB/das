package competition

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/controller/competition"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiCompetitionStatusEndpoint = "/api/competition/status"

var competitionStatusServer = competition.CompetitionStatusServer{
	database.CompetitionStatusRepository,
}

var GetCompetitionStatusController = util.DasController{
	Name:         "GetCompetitionStatusController",
	Description:  "Get all competition status",
	Method:       http.MethodGet,
	Endpoint:     apiCompetitionStatusEndpoint,
	Handler:      competitionStatusServer.GetCompetitionStatusHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_NOAUTH},
}
