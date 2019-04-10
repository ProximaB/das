package dalutil

import (
	"fmt"
	"reflect"
)

// DataSourceNotSpecifiedError takes a Repository object and return a generic error message
func DataSourceNotSpecifiedError(repo interface{}) string {
	return fmt.Sprintf("data source of %s is not specified", reflect.TypeOf(repo).String())
}

const ErrorNilDatabase = "should throw an error when repository is not initialized correctly"
