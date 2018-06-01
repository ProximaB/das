package main

import (
	"github.com/yubing24/das/config/database"
	"github.com/yubing24/das/config/routes"
	"github.com/yubing24/das/controller/util"
	"google.golang.org/appengine"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func decodeIntArray(value string) reflect.Value {
	value = strings.Replace(value, "[", "", -1)
	value = strings.Replace(value, "]", "", -1)
	split := strings.Split(value, ",")
	data := make([]int, 0)
	for _, each := range split {
		val, err := strconv.Atoi(each)
		if err != nil {
			return reflect.Value{}
		}
		data = append(data, val)
	}
	return reflect.ValueOf(data)
}

// register this to gorilla schema decoder to decode time
func decodeDate(value string) reflect.Value {
	s, err := time.Parse("2006-01-02", value)
	if err != nil {
		return reflect.ValueOf(DasParseHtmlInputDateTime(value))
	}
	return reflect.ValueOf(s)
}

func DasParseHtmlInputDateTime(input string) time.Time {
	layout := "2006-01-02T15:04"
	t, _ := time.Parse(layout, input)
	return t
}

func DasParseHtmlInputDate(input string) time.Time {
	layout := "2006-01-02"
	t, _ := time.Parse(layout, input)
	return t
}

func configureController() {
	util.JsonRequestDecoder.RegisterConverter(time.Time{}, decodeDate)
	util.JsonRequestDecoder.RegisterConverter([]int{}, decodeIntArray)

	/*controller.HMAC_SECRET = os.Getenv("HMAC_SIGNING_KEY")
	if len(controller.HMAC_SECRET) == 0 {
		// use a default key, not recommended!!!
		controller.HMAC_SECRET = "DAS DEFAULT SIGNING KEY"
	}*/
}

func main() {
	configureController()

	defer database.PostgresDatabase.Close() // database connection will not close until server is shutdown
	router := routes.DasRouter()

	if database.PostgresDatabase == nil {
		log.Fatal("cannot establish connection to provided database")
	}

	http.Handle("/", router)
	appengine.Main() // to run this on app engine, do not make router listen to any particular port
}
