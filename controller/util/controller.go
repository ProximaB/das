// Dancesport Application System (DAS)
// Copyright (C) 2017, 2018 Yubing Hou
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

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
