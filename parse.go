package nut

import (
	"reflect"
	"strings"
)

func parse(data any) (*Structure, error) {
	rt := reflect.TypeOf(data)
	//rv := reflect.ValueOf(data)

	opt := &Structure{
		StructFullName:  "",
		StructShortName: "",
	}

	n := rt.NumField()
	for i := 0; i < n; i++ {
		field := rt.Field(i)
		tv := field.Tag.Get(Nut)

		cns := strings.Split(tv, ";")
		err := check(field, cns)
		if err != nil {
			return nil, err
		}
	}

	return opt, nil
}
