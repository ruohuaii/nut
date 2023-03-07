package nut

import (
	"fmt"
	"reflect"
	"strings"
)

func parse(data any) (specimen, error) {
	opt := specimen{}
	rt := reflect.TypeOf(data)
	if rt.Kind() == reflect.Pointer {
		rt = rt.Elem()
	}

	//set FullName,shortName
	opt.FullName = fmt.Sprintf("*%s", rt.Name())
	opt.ShortName = strings.ToLower(rt.Name()[:1])

	associateRules := make(map[string][]Relation)
	conditions := make(map[string][]Condition)
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
				case Associate:
					associateRules[fieldName] = append(associateRules[fieldName], Relation{
						Self:      fieldName,
						Associate: strings.Split(fcv[1], ","),
					})
				case Eq:
					conditions[fieldName] = append(conditions[fieldName], Condition{
						Description: ThrowCondNeq(opt.ShortName, fieldName, fcv[1]),
					})
				case Neq:
					conditions[fieldName] = append(conditions[fieldName], Condition{
						Description: ThrowCondEq(opt.ShortName, fieldName, fcv[1]),
					})
				case Lt:
					conditions[fieldName] = append(conditions[fieldName], Condition{
						Description: ThrowCondGte(opt.ShortName, fieldName, fcv[1]),
					})
				case Lte:
					conditions[fieldName] = append(conditions[fieldName], Condition{
						Description: ThrowCondGt(opt.ShortName, fieldName, fcv[1]),
					})
				case Gt:
					conditions[fieldName] = append(conditions[fieldName], Condition{
						Description: ThrowCondLte(opt.ShortName, fieldName, fcv[1]),
					})
				case Gte:
					conditions[fieldName] = append(conditions[fieldName], Condition{
						Description: ThrowCondLt(opt.ShortName, fieldName, fcv[1]),
					})
				case Between:
					cvs := strings.Split(fcv[1], ",")
					conditions[fieldName] = append(conditions[fieldName], Condition{
						Description: ThrowCondBetween(opt.ShortName, fieldName, cvs[0], cvs[1]),
					})
				case Size:
					cvs := strings.Split(fcv[1], ",")
					conditions[fieldName] = append(conditions[fieldName], Condition{
						Description: ThrowCondSize(opt.ShortName, fieldName, cvs, kind.String()),
					})
				case Regexp:
					conditions[fieldName] = append(conditions[fieldName], Condition{
						Description: ThrowCondRegexp(opt.ShortName, fieldName, fcv[1]),
					})
				case In:
					var elemType string
					for _, v := range fcs {
						rule := strings.Split(v, ":")
						if rule[0] == Type {
							elemType = rule[1]
						}
					}
					conditions[fieldName] = append(conditions[fieldName], Condition{
						Description: ThrowCondIn(opt.ShortName, fieldName, fcv[1], elemType),
					})
				}
			case reflect.Slice, reflect.Array:
				var elemType string
				for _, v := range fcs {
					rule := strings.Split(v, ":")
					if rule[0] == Type {
						elemType = rule[1]
					}
				}
				switch fcv[0] {
				case Contains:
					conditions[fieldName] = append(conditions[fieldName], Condition{
						Description: ThrowCondExcluded(opt.ShortName, fieldName, fcv[1], elemType),
					})
				case Excluded:
					conditions[fieldName] = append(conditions[fieldName], Condition{
						Description: ThrowCondContains(opt.ShortName, fieldName, fcv[1], elemType),
					})
				case Size:
					cvs := strings.Split(fcv[1], ",")
					conditions[fieldName] = append(conditions[fieldName], Condition{
						Description: ThrowCondSize(opt.ShortName, fieldName, cvs, kind.String()),
					})
				}
			case reflect.Struct:

			}

		}
	}

	opt.Associates = associateRules
	opt.Conditions = conditions

	return opt, nil
}
