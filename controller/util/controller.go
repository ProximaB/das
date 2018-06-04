// a new idea for managing endpoints in DAS. this is still in experiment and should not be added to production until fully
// tested
package util

import (
	"net/http"
)

// DasController specifies the description, allowed method, endpoint, handler functions,
// and roles allowed to access this controller. Controller does not specifies the underlying
// source of data that it depends on. Instead, create a separate sever for each particular
// controller, specify the data source inside the server struct and inject data source into
// controller's HandlerFunc implementation.
type DasController struct {
	Name         string
	Description  string
	Method       string
	Endpoint     string
	Handler      http.HandlerFunc
	AllowedRoles []int
}

type DasControllerGroup struct {
	Controllers []DasController
}
