package nut

import (
	"fmt"
	"reflect"
	"strings"
)

func parse(rt reflect.Type) (specimen, error) {
	opt := specimen{}
	if rt.Kind() == reflect.Pointer {
		rt = rt.Elem()
	}

	//set FullName,shortName
	opt.FullName = fmt.Sprintf("*%s", rt.Name())
	opt.ShortName = strings.ToLower(rt.Name()[:1])

	relationRules := make(map[string]Relation)
	selfRules := make(map[string][]Condition)
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
					relationRules[fieldName] = Relation{
						Self:      fieldName,
						Associate: fcv[1],
					}
				case Eq:
					selfRules[fieldName] = append(selfRules[fieldName], Condition{
						Rule: ThrowCondNeq(opt.ShortName, fieldName, fcv[1]),
					})
				case Neq:
					selfRules[fieldName] = append(selfRules[fieldName], Condition{
						Rule: ThrowCondEq(opt.ShortName, fieldName, fcv[1]),
					})
				case Lt:
					selfRules[fieldName] = append(selfRules[fieldName], Condition{
						Rule: ThrowCondGte(opt.ShortName, fieldName, fcv[1]),
					})
				case Lte:
					selfRules[fieldName] = append(selfRules[fieldName], Condition{
						Rule: ThrowCondGt(opt.ShortName, fieldName, fcv[1]),
					})
				case Gt:
					selfRules[fieldName] = append(selfRules[fieldName], Condition{
						Rule: ThrowCondLte(opt.ShortName, fieldName, fcv[1]),
					})
				case Gte:
					selfRules[fieldName] = append(selfRules[fieldName], Condition{
						Rule: ThrowCondLt(opt.ShortName, fieldName, fcv[1]),
					})
				case Between:
					cvs := strings.Split(fcv[1], ",")
					selfRules[fieldName] = append(selfRules[fieldName], Condition{
						Rule: ThrowCondBetween(opt.ShortName, fieldName, cvs[0], cvs[1]),
					})
				case Size:
					cvs := strings.Split(fcv[1], ",")
					selfRules[fieldName] = append(selfRules[fieldName], Condition{
						Rule: ThrowCondSize(opt.ShortName, fieldName, cvs, kind.String()),
					})
				case Regexp:
					selfRules[fieldName] = append(selfRules[fieldName], Condition{
						Rule: ThrowCondRegexp(opt.ShortName, fieldName, fcv[1]),
					})
				case In:
					selfRules[fieldName] = append(selfRules[fieldName], Condition{
						Rule: ThrowCondIn(opt.ShortName, fieldName, fcv[1], kind.String()),
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
					selfRules[fieldName] = append(selfRules[fieldName], Condition{
						Rule: ThrowCondExcluded(opt.ShortName, fieldName, fcv[1], elemType),
					})
				case Excluded:
					selfRules[fieldName] = append(selfRules[fieldName], Condition{
						Rule: ThrowCondContains(opt.ShortName, fieldName, fcv[1], elemType),
					})
				case Size:
					cvs := strings.Split(fcv[1], ",")
					selfRules[fieldName] = append(selfRules[fieldName], Condition{
						Rule: ThrowCondSize(opt.ShortName, fieldName, cvs, kind.String()),
					})
				}
			}

		}
	}

	associateRules := make(map[string][]Condition)
	for k := range selfRules {
		if m, ok := relationRules[k]; ok {
			associateRules[m.Associate] = append(associateRules[m.Associate], selfRules[k]...)
			delete(selfRules, k)
		}
	}

	conditions := make(map[string]Rules)

	for k, v := range selfRules {
		conditions[k] = Rules{
			SelfRules:      v,
			AssociateRules: associateRules[k],
		}
	}

	opt.Conditions = conditions

	return opt, nil
}
