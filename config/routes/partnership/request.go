package partnership

import (
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/config/database"
	"github.com/ProximaB/das/config/routes/middleware"
	"github.com/ProximaB/das/controller/partnership/request"
	"github.com/ProximaB/das/controller/util"
	"net/http"
)

const apiPartnershipRequestEndpoint = "/api/v1.0/athlete/partnership/request"

var partnershipRequestServer = request.PartnershipRequestServer{
	middleware.AuthenticationStrategy,
	database.AccountRepository,
	database.PartnershipRepository,
	database.PartnershipRequestRepository,
	database.PartnershipRequestBlacklistRepository,
}

var createPartnershipRequestController = util.DasController{
	Name:         "CreatePartnershipRequestController",
	Description:  "Create a new partnership request in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiPartnershipRequestEndpoint,
	Handler:      partnershipRequestServer.CreatePartnershipRequestHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAthlete},
}

var searchPartnershipRequestController = util.DasController{
	Name:         "SearchPartnershipRequestController",
	Description:  "Search a new partnership request in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiPartnershipRequestEndpoint,
	Handler:      partnershipRequestServer.SearchPartnershipRequestHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAthlete},
}

var updatePartnershipRequestController = util.DasController{
	Name:         "UpdatePartnershipRequestController",
	Description:  "Update a new partnership request in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiPartnershipRequestEndpoint,
	Handler:      partnershipRequestServer.UpdatePartnershipRequestHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAthlete},
}

var deletePartnershipRequestController = util.DasController{
	Name:         "DeletePartnershipRequestController",
	Description:  "delete a new partnership request in DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiPartnershipRequestEndpoint,
	Handler:      partnershipRequestServer.DeletePartnershipRequestHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAthlete},
}

var PartnershipRequestControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		createPartnershipRequestController,
		searchPartnershipRequestController,
		updatePartnershipRequestController,
		deletePartnershipRequestController,
	},
}
