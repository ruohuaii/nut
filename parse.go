package nut

import (
	"reflect"
)

func parse(data any) (*Structure, error) {
	//rt := reflect.TypeOf(data)
	rv := reflect.ValueOf(data)

	opt := &Structure{
		StructFullName:  "",
		StructShortName: "",
	}

	err := depthCheck(rv)
	if err != nil {
		return nil, err
	}

	return opt, nil
}
