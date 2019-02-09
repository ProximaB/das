package account

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/controller/account"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiAccountGenderEndpoint = "/api/account/gender"

var genderServer = account.GenderServer{
	database.GenderRepository,
}

var GenderController = util.DasController{
	Name:         "GenderController",
	Description:  "Get all genders in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiAccountGenderEndpoint,
	Handler:      genderServer.GetAccountGenderHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}
