package partnership

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/config/database"
	"github.com/yubing24/das/controller/partnership"
	"github.com/yubing24/das/controller/util"
	"net/http"
)

const apiPartnershipEndpoint = "/api/partnership"

var partnershipServer = partnership.PartnershipServer{
	database.AccountRepository, database.PartnershipRepository,
}

var searchPartnershipController = util.DasController{
	Description:  "Search partnerships in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiPartnershipEndpoint,
	Handler:      partnershipServer.SearchPartnershipHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ATHLETE},
}

var updatePartnershipController = util.DasController{
	Description:  "Update a partnership in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiPartnershipEndpoint,
	Handler:      partnershipServer.UpdatePartnershipHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ATHLETE},
}

var PartnershipControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchPartnershipController,
		updatePartnershipController,
	},
}
