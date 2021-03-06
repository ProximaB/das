package account

import (
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/config/database"
	"github.com/ProximaB/das/controller/account"
	"github.com/ProximaB/das/controller/util"
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
