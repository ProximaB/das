package competition

import (
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/config/database"
	"github.com/ProximaB/das/controller/competition"
	"github.com/ProximaB/das/controller/util"
	"net/http"
)

const apiCompetitionStatusEndpoint = "/api/competition/status"

var competitionStatusServer = competition.StatusServer{
	database.CompetitionStatusRepository,
}

var GetCompetitionStatusController = util.DasController{
	Name:         "GetCompetitionStatusController",
	Description:  "Get all competition status",
	Method:       http.MethodGet,
	Endpoint:     apiCompetitionStatusEndpoint,
	Handler:      competitionStatusServer.GetStatusHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}
