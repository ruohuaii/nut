package nut

import "reflect"

func Generate(data any) (*Structure, error) {
	rv := reflect.ValueOf(data)
	err := depthCheck(rv)
	if err != nil {
		return nil, err
	}

	opt, _ := parse(data)

	err = FillTemplate(tpl, opt, "main.go")
	if err != nil {
		return nil, err
	}

	return &opt, nil
}
