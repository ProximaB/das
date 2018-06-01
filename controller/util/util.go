package util

import (
	"github.com/yubing24/das/viewmodel"
	"encoding/json"
	"github.com/gorilla/schema"
	"net/http"
)

const (
	HTTP_400_INVALID_REQUEST_DATA  = "invalid request data"
	HTTP_500_ERROR_RETRIEVING_DATA = "error in retrieving data"
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
