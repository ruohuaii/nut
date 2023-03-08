package nut

import (
	"reflect"
	"strings"
)

func Generate(data any, file string) error {
	mainValue := reflect.ValueOf(data)
	if mainValue.Kind() == reflect.Pointer && !mainValue.IsNil() {
		mainValue = mainValue.Elem()
	}

	err := depthCheck(mainValue)
	if err != nil {
		return err
	}

	mainType := reflect.TypeOf(data)
	if mainType.Kind() == reflect.Pointer {
		mainType = mainType.Elem()
	}
	types := pickStruct(mainType)
	types = append(types, FieldStruct{
		Type: mainType,
	})

	mainShortName := strings.ToLower(mainType.Name()[:1])
	structFieldCondSet := make([]Condition, 0)
	last := len(types) - 1
	for i := 0; i < len(types); i++ {
		isStructField := false
		if types[i].FieldName != "" {
			isStructField = true
		}
		opt, structFieldCond, err := parse(
			types[i].Type, isStructField, types[i].FieldName, mainShortName,
			types[i].IsOptional, types[i].IsPtr,
		)
		if err != nil {
			return err
		}
		if i == last {
			opt.Conditions[StructField] = Rules{
				SelfRules: structFieldCondSet,
			}
		} else {
			structFieldCondSet = append(structFieldCondSet, structFieldCond)
		}
		err = FillTemplate(tpl, opt, file)
		if err != nil {
			return err
		}
	}

	return nil
}
