package nut

import (
	"reflect"
)

func Generate(data any, file string) (*specimen, error) {
	rv := reflect.ValueOf(data)
	err := depthCheck(rv)
	if err != nil {
		return nil, err
	}

	opt, err := parse(reflect.TypeOf(data))
	if err != nil {
		return nil, err
	}

	err = FillTemplate(tpl, opt, file)
	if err != nil {
		return nil, err
	}

	return &opt, nil
}
