package account

import (
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/config/database"
	"github.com/ProximaB/das/controller/account"
	"github.com/ProximaB/das/controller/util"
	"net/http"
)

var profileSearchServer = account.NewProfileSearchServer(database.AccountRepository)

var searchDancerProfileController = util.DasController{
	Name:         "SearchDancerProfileController",
	Description:  "Search Dancers' Profiles",
	Method:       http.MethodGet,
	Endpoint:     "/api/v1.0/profile/dancer",
	Handler:      profileSearchServer.SearchDancerProfileHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}
var searchPartnershipProfileController = util.DasController{
	Name:         "SearchPartnershipProfileController",
	Description:  "Search Partnerships' Profiles",
	Method:       http.MethodGet,
	Endpoint:     "/api/v1.0/profile/partnership",
	Handler:      profileSearchServer.SearchPartnershipProfileHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

var SearchProfileControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchPartnershipProfileController,
		searchDancerProfileController,
	},
}
