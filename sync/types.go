package sync

import "reflect"

func IsString(v interface{}) bool {

	return reflect.TypeOf(v).Kind() == reflect.String
}
