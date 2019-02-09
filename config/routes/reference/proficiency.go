package reference

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/controller/reference"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiReferenceProficiencyEndpoint = "/api/v1.0/reference/proficiency"

var proficiencyServer = reference.ProficiencyServer{
	database.ProficiencyRepository,
}

var searchProficiencyController = util.DasController{
	Name:         "SearchProficiencyController",
	Description:  "Search proficiencies in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiReferenceProficiencyEndpoint,
	Handler:      proficiencyServer.SearchProficiencyHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

var createProficiencyController = util.DasController{
	Name:         "CreateProficiencyController",
	Description:  "Create a proficiency in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiReferenceProficiencyEndpoint,
	Handler:      proficiencyServer.CreateProficiencyHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var deleteProficiencyController = util.DasController{
	Name:         "DeleteProficiencyController",
	Description:  "Delete a proficiency from DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiReferenceProficiencyEndpoint,
	Handler:      proficiencyServer.DeleteProficiencyHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var updateProficiencyController = util.DasController{
	Name:         "UpdateProficiencyController",
	Description:  "Update a proficiency in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiReferenceProficiencyEndpoint,
	Handler:      proficiencyServer.UpdateProficiencyHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

// ProficiencyControllerGroup is a collection of handler functions for managing proficiencies in DAS
var ProficiencyControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchProficiencyController,
		createProficiencyController,
		deleteProficiencyController,
		updateProficiencyController,
	},
}
