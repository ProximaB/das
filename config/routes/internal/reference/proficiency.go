package reference

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/config/database"
	"github.com/yubing24/das/controller/reference"
	"github.com/yubing24/das/controller/util"
	"net/http"
)

const apiReferenceProficiencyEndpoint = "/api/reference/proficiency"

var proficiencyServer = reference.ProficiencyServer{
	database.ProficiencyRepository,
}

var searchProficiencyController = util.DasController{
	Description:  "Search proficiencies in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiReferenceProficiencyEndpoint,
	Handler:      proficiencyServer.SearchProficiencyHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_NOAUTH},
}

var createProficiencyController = util.DasController{
	Description:  "Create a proficiency in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiReferenceProficiencyEndpoint,
	Handler:      proficiencyServer.CreateProficiencyHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var deleteProficiencyController = util.DasController{
	Description:  "Delete a proficiency from DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiReferenceProficiencyEndpoint,
	Handler:      proficiencyServer.DeleteProficiencyHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var updateProficiencyController = util.DasController{
	Description:  "Update a proficiency in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiReferenceProficiencyEndpoint,
	Handler:      proficiencyServer.UpdateProficiencyHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var ProficiencyControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchProficiencyController,
		createProficiencyController,
		deleteProficiencyController,
		updateProficiencyController,
	},
}
