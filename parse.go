package nut

import (
	"fmt"
	"reflect"
	"strings"
)

func parse(rt reflect.Type, isStructField bool, fieldName string, mainShortName string, isOptional, isPtr bool) (
	*specimen, Condition, error) {
	opt := &specimen{}
	if rt.Kind() == reflect.Pointer {
		rt = rt.Elem()
	}

	//set FullName,shortName
	opt.FullName = fmt.Sprintf("*%s", rt.Name())
	opt.ShortName = strings.ToLower(rt.Name()[:1])

	relationRules := make(map[string]Relation)
	selfRules := make(map[string]map[string]Condition)
	var structFieldCond Condition
	for i := 0; i < rt.NumField(); i++ {
		fieldName := rt.Field(i).Name
		tag := rt.Field(i).Tag.Get(Nut)
		definedRules := getDefinedRules(tag)
		for k, v := range definedRules {
			if _, ok := selfRules[fieldName]; !ok {
				selfRules[fieldName] = make(map[string]Condition)
			}
			kind := rt.Field(i).Type.Kind()
			if ArrayContains(SpecialCondSet[:], k) {
				selfRules[fieldName][k] = Condition{
					Rule:  k,
					FType: kind.String(),
				}
				continue
			}
			switch kind {
			case reflect.Int8, reflect.Int16, reflect.Int32,
				reflect.Int, reflect.Int64, reflect.Uint8,
				reflect.Uint16, reflect.Uint32, reflect.Uint,
				reflect.Uint64, reflect.Float32, reflect.Float64,
				reflect.String:
				switch k {
				case Associate:
					relationRules[fieldName] = Relation{
						Self:      fieldName,
						Associate: k,
					}
				case Eq:
					isString := false
					if kind == reflect.String {
						isString = true
					}
					selfRules[fieldName][Eq] = Condition{
						Rule: ThrowCondNeq(opt.ShortName, fieldName, v, isString),
					}
				case Neq:
					isString := false
					if kind == reflect.String {
						isString = true
					}
					selfRules[fieldName][Neq] = Condition{
						Rule: ThrowCondEq(opt.ShortName, fieldName, v, isString),
					}
				case Lt:
					selfRules[fieldName][Lt] = Condition{
						Rule: ThrowCondGte(opt.ShortName, fieldName, v),
					}
				case Lte:
					selfRules[fieldName][Lte] = Condition{
						Rule: ThrowCondGt(opt.ShortName, fieldName, v),
					}
				case Gt:
					selfRules[fieldName][Gt] = Condition{
						Rule: ThrowCondLte(opt.ShortName, fieldName, v),
					}
				case Gte:
					selfRules[fieldName][Gte] = Condition{
						Rule: ThrowCondLt(opt.ShortName, fieldName, v),
					}
				case Between:
					cvs := strings.Split(v, ",")
					selfRules[fieldName][Between] = Condition{
						Rule: ThrowCondBetween(opt.ShortName, fieldName, cvs[0], cvs[1]),
					}
				case Size:
					cvs := strings.Split(v, ",")
					selfRules[fieldName][Size] = Condition{
						Rule: ThrowCondSize(opt.ShortName, fieldName, cvs, kind.String()),
					}
				case Regexp:
					selfRules[fieldName][Regexp] = Condition{
						Rule: ThrowCondRegexp(opt.ShortName, fieldName, v),
					}
				case In:
					selfRules[fieldName][In] = Condition{
						Rule: ThrowCondIn(opt.ShortName, fieldName, v, kind.String()),
					}
				case Type:
					elemType := definedRules[Type]
					selfRules[fieldName][Type] = Condition{
						Rule: ThrowCondType(opt.ShortName, fieldName, elemType),
					}
				}
			case reflect.Slice, reflect.Array:
				elemType := definedRules[Type]
				switch k {
				case Contains:
					selfRules[fieldName][Contains] = Condition{
						Rule: ThrowCondExcluded(opt.ShortName, fieldName, v, elemType),
					}
				case Excluded:
					selfRules[fieldName][Excluded] = Condition{
						Rule: ThrowCondContains(opt.ShortName, fieldName, v, elemType),
					}
				case Size:
					cvs := strings.Split(v, ",")
					selfRules[fieldName][Size] = Condition{
						Rule: ThrowCondSize(opt.ShortName, fieldName, cvs, kind.String()),
					}
				}
			}
		}
	}

	if isStructField {
		structFieldCond = Condition{
			Rule: ThrowCondStruct(mainShortName, fieldName, rt.Name(), isOptional, isPtr),
		}
	}

	//Association conditions will be implemented in the future
	_ = relationRules
	conditions := make(map[string]Rules)

	for f, v := range selfRules {
		rules := Rules{}
		rules.SelfRules = make([]Condition, 0)
		for c, m := range v {
			if c == Optional {
				switch m.FType {
				case String:
					format1 := `if %s.%s != ""{`
					topRule := fmt.Sprintf(format1, opt.ShortName, f)
					for _, r := range v {
						if r.Rule == Optional || r.Rule == Required {
							continue
						}
						topRule += r.Rule + "\n"
					}
					rule := topRule + "}"
					rules.SelfRules = append(rules.SelfRules, Condition{
						Rule: rule,
					})

				case Int8, Int16, Int32, Int, Int64, Uint8, Uint16, Uint32, Uint, Uint64, Float32, Float64:
					format1 := `if %s.%s != 0{`
					topRule := fmt.Sprintf(format1, opt.ShortName, f)
					for _, r := range v {
						topRule += r.Rule + "\n"
					}
					rule := topRule + "}"
					rules.SelfRules = append(rules.SelfRules, Condition{
						Rule: rule,
					})
				}
			} else {
				if c == Required {
					continue
				}
				rules.SelfRules = append(rules.SelfRules, m)
			}
		}
		conditions[f] = rules
	}

	opt.Conditions = conditions

	return opt, structFieldCond, nil
}

func pickStruct(mainType reflect.Type) []FieldStruct {
	types := make([]FieldStruct, 0)
	for i := 0; i < mainType.NumField(); i++ {
		fieldType := mainType.Field(i).Type
		switch fieldType.Kind() {
		case reflect.Struct:
			types = append(types, FieldStruct{
				FieldName:  mainType.Field(i).Name,
				Type:       fieldType,
				IsOptional: strings.Contains(mainType.Field(i).Tag.Get(Nut), Optional),
				IsPtr:      false,
			})
		case reflect.Pointer:
			fieldType := mainType.Field(i).Type.Elem()
			if fieldType.Kind() == reflect.Struct {
				types = append(types, FieldStruct{
					FieldName:  mainType.Field(i).Name,
					Type:       fieldType,
					IsOptional: strings.Contains(mainType.Field(i).Tag.Get(Nut), Optional),
					IsPtr:      true,
				})
			}
		}
	}

	return types
}

func getDefinedRules(tag string) map[string]string {
	rules := strings.Split(tag, Semicolon)
	definedRules := make(map[string]string)
	for _, v := range rules {
		conditions := strings.Split(v, ":")
		if len(conditions) == 0 {
			continue
		} else if len(conditions) == 1 {
			definedRules[conditions[0]] = StringNull
		} else {
			definedRules[conditions[0]] = conditions[1]
		}
	}

	return definedRules
}
