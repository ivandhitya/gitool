package structs

import (
	"fmt"
	"reflect"
)

const (
	DEFAULT_TAG = "mapkey"
)

func StructToMapString(structs interface{}) (map[string]string, error) {
	m := make(map[string]string)
	val := reflect.ValueOf(structs)
	for i := 0; i < val.Type().NumField(); i++ {
		valTemp := reflect.ValueOf(val.Field(i).Interface()).Interface()
		if val.Field(i).Kind() == reflect.Ptr && val.Field(i).IsNil() {
			continue
		}
		key := val.Type().Field(i).Tag.Get(DEFAULT_TAG)
		value := fmt.Sprintf("%v", valTemp)
		m[key] = value
	}
	return m, nil
}
