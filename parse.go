package nut

import (
	"fmt"
	"reflect"
	"strings"
)

func parse(data any) (*Structure, error) {
	rv := reflect.ValueOf(data)
	err := depthCheck(rv)
	if err != nil {
		return nil, err
	}

	rt := reflect.TypeOf(data)
	opt := Structure{
		FullName:  fmt.Sprintf("*%s", rt.Name()),
		ShortName: strings.ToLower(rt.Name()[:1]),
	}

	err = FillTemplate(tpl, opt, "main.go")
	if err != nil {
		return nil, err
	}

	return &opt, nil
}
