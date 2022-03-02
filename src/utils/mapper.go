package utils

import (
	"reflect"

	"github.com/thitiratratrat/hhor/src/customtype"
)

func Map(s interface{}) map[string]interface{} {
	sMap := make(map[string]interface{})

	v := reflect.ValueOf(s)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		value, err := v.Field(i).Interface().(customtype.Nullable).GetValue()

		if err != nil {
			continue
		}

		sMap[typeOfS.Field(i).Name] = value
	}

	return sMap
}
