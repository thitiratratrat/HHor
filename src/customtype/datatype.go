package customtype

import (
	"encoding/json"
	"errors"
)

type Nullable interface {
	GetValue() (interface{}, error)
}

type JSONString struct {
	Value  string
	IsNull bool
	Set    bool
}

type JSONInt struct {
	Value  int
	IsNull bool
	Set    bool
}

func (i *JSONString) UnmarshalJSON(data []byte) error {
	i.Set = true

	if string(data) == "null" {
		i.IsNull = true
		return nil
	}

	var temp string
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	i.Value = temp
	i.IsNull = false
	return nil
}

func (i JSONString) GetValue() (interface{}, error) {
	if !i.Set {
		return nil, errors.New("JSON field not set")
	}

	if i.IsNull {
		return nil, nil
	}

	return i.Value, nil
}

func (i *JSONInt) UnmarshalJSON(data []byte) error {
	i.Set = true

	if string(data) == "null" {
		i.IsNull = true
		return nil
	}

	var temp int
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	i.Value = temp
	i.IsNull = false
	return nil
}

func (i JSONInt) GetValue() (interface{}, error) {
	if !i.Set {
		return nil, errors.New("JSON field not set")
	}

	if i.IsNull {
		return nil, nil
	}

	return i.Value, nil
}
