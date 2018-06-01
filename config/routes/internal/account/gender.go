package account

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/config/database"
	"github.com/yubing24/das/controller/account"
	"github.com/yubing24/das/controller/util"
	"net/http"
)

const apiAccountGenderEndpoint = "/api/account/gender"

var genderServer = account.GenderServer{
	database.GenderRepository,
}

var GenderController = util.DasController{
	Description:  "Get all genders in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiAccountGenderEndpoint,
	Handler:      genderServer.GetAccountGenderHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_NOAUTH},
}
