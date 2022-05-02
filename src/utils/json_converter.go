package utils

import "encoding/json"

func ToJson(val []byte) interface{} {
	var obj interface{}
	err := json.Unmarshal(val, &obj)

	if err != nil {
		panic(err)
	}

	return obj
}
