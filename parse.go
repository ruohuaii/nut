package nut

import (
	"fmt"
	"reflect"
	"strings"
)

func parse(data any) (Structure, error) {
	opt := Structure{}
	rt := reflect.TypeOf(data)

	//set FullName,ShortName
	opt.FullName = fmt.Sprintf("*%s", rt.Name())
	opt.ShortName = strings.ToLower(rt.Name()[:1])

	rules := make(map[string][]Field)
	for i := 0; i < rt.NumField(); i++ {
		fieldName := rt.Field(i).Name
		tag := rt.Field(i).Tag.Get(Nut)
		fcs := strings.Split(tag, ";")
		for _, v := range fcs {
			fcv := strings.Split(v, ":")
			if len(fcv) != 2 {
				continue
			}
			//fcv[0] is the nut condition type
			kind := rt.Field(i).Type.Kind()
			switch kind {
			case reflect.Int8, reflect.Int16, reflect.Int32,
				reflect.Int, reflect.Int64, reflect.Uint8,
				reflect.Uint16, reflect.Uint32, reflect.Uint,
				reflect.Uint64, reflect.Float32, reflect.Float64,
				reflect.String:
				switch fcv[0] {
				case Eq:
					rules[fieldName] = append(rules[fieldName], Field{
						Condition: ThrowCondNeq(opt.ShortName, fieldName, fcv[1]),
					})
				case Neq:
					rules[fieldName] = append(rules[fieldName], Field{
						Condition: ThrowCondEq(opt.ShortName, fieldName, fcv[1]),
					})
				case Lt:
					rules[fieldName] = append(rules[fieldName], Field{
						Condition: ThrowCondLt(opt.ShortName, fieldName, fcv[1]),
					})
				case Lte:
					rules[fieldName] = append(rules[fieldName], Field{
						Condition: ThrowCondLte(opt.ShortName, fieldName, fcv[1]),
					})
				case Gt:
					rules[fieldName] = append(rules[fieldName], Field{
						Condition: ThrowCondGt(opt.ShortName, fieldName, fcv[1]),
					})
				case Gte:
					rules[fieldName] = append(rules[fieldName], Field{
						Condition: ThrowCondGte(opt.ShortName, fieldName, fcv[1]),
					})
				case Between:
					cvs := strings.Split(fcv[1], ",")
					rules[fieldName] = append(rules[fieldName], Field{
						Condition: ThrowCondBetween(opt.ShortName, fieldName, cvs[0], cvs[1]),
					})
				case Size:
					cvs := strings.Split(fcv[1], ",")
					rules[fieldName] = append(rules[fieldName], Field{
						Condition: ThrowCondSize(opt.ShortName, fieldName, cvs[0], cvs[1], kind.String()),
					})
				}

			}

		}
	}

	opt.Rules = rules

	return opt, nil
}
