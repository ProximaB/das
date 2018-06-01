package partnership

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/config/database"
	"github.com/yubing24/das/controller/partnership/request"
	"github.com/yubing24/das/controller/util"
	"net/http"
)

const apiPartnershipRequestEndpoint = "/api/partnership/request"

var partnershipRequestServer = request.PartnershipRequestServer{
	database.AccountRepository,
	database.PartnershipRepository,
	database.PartnershipRequestRepository,
	database.PartnershipRequestBlacklistRepository,
}

var createPartnershipRequestController = util.DasController{
	Description:  "Create a new partnership request in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiPartnershipRequestEndpoint,
	Handler:      partnershipRequestServer.CreatePartnershipRequestHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ATHLETE},
}

var searchPartnershipRequestController = util.DasController{
	Description:  "Search a new partnership request in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiPartnershipRequestEndpoint,
	Handler:      partnershipRequestServer.SearchPartnershipRequestHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ATHLETE},
}

var updatePartnershipRequestController = util.DasController{
	Description:  "Update a new partnership request in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiPartnershipRequestEndpoint,
	Handler:      partnershipRequestServer.UpdatePartnershipRequestHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ATHLETE},
}

var deletePartnershipRequestController = util.DasController{
	Description:  "delete a new partnership request in DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiPartnershipRequestEndpoint,
	Handler:      partnershipRequestServer.DeletePartnershipRequestHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ATHLETE},
}

var PartnershipRequestControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		createPartnershipRequestController,
		searchPartnershipRequestController,
		updatePartnershipRequestController,
		deletePartnershipRequestController,
	},
}
