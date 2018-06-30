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
	"encoding/json"
	"github.com/DancesportSoftware/das/viewmodel"
	"github.com/gorilla/schema"
	"net/http"
)

const (
	// HTTP400InvalidRequestData provides a generic message that can be used when HTTP 400 error has to be returned
	HTTP400InvalidRequestData = "invalid request data"
	// HTTP500ErrorRetrievingData provides a generic error message which indicates that data access layer code has thrown an error
	HTTP500ErrorRetrievingData = "error in retrieving data"
)

var JsonRequestDecoder = schema.NewDecoder()

func ParseRequestData(r *http.Request, dto interface{}) error {
	r.ParseForm()
	err := JsonRequestDecoder.Decode(dto, r.Form)
	return err
}

func ParseRequestBodyData(r *http.Request, dto interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(&dto)
}

func RespondJsonResult(w http.ResponseWriter, status int, message string, data interface{}) {
	result := viewmodel.RESTAPIResult{
		Status:  status,
		Message: message,
		Data:    data,
	}
	output, _ := json.Marshal(result)
	//w.WriteHeader(status) // this affects the behavior on angular
	w.Write(output)
}
