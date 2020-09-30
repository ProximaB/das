package partnership

import (
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/config/database"
	"github.com/ProximaB/das/config/routes/middleware"
	"github.com/ProximaB/das/controller/partnership"
	"github.com/ProximaB/das/controller/util"
	"net/http"
)

const apiPartnershipEndpoint = "/api/v1.0/athlete/partnership"

var partnershipServer = partnership.PartnershipServer{
	middleware.AuthenticationStrategy,
	database.AccountRepository,
	database.PartnershipRepository,
}

var searchPartnershipController = util.DasController{
	Name:         "SearchPartnershipController",
	Description:  "Search partnerships in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiPartnershipEndpoint,
	Handler:      partnershipServer.SearchPartnershipHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAthlete},
}

var updatePartnershipController = util.DasController{
	Name:         "UpdatePartnershipController",
	Description:  "Update a partnership in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiPartnershipEndpoint,
	Handler:      partnershipServer.UpdatePartnershipHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAthlete},
}

// PartnershipControllerGroup contains a collection of HTTP request handler functions for
// Partnership related request
var PartnershipControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchPartnershipController,
		updatePartnershipController,
	},
}
