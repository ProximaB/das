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
	Http404NoDataFound        = "no data is found"
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
	w.WriteHeader(status)
	w.Write(output)
}
