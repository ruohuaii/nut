package nut

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func depthCheck(v reflect.Value) error {
	if !v.CanInterface() {
		return errors.New("exists on an unsafe type")
	}

	err := checkRule(v)
	if err != nil {
		return err
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fv := v.Field(i)
		switch fv.Kind() {
		case reflect.Struct:
			fieldName := t.Field(i).Name
			tag := t.Field(i).Tag.Get(Nut)
			err := checkStruct(fv, fieldName, tag)
			if err != nil {
				return err
			}
		case reflect.Pointer:
			if fv.IsNil() {
				continue
			}
			fvc := fv.Elem()
			switch fvc.Kind() {
			case reflect.Struct:
				fieldName := t.Field(i).Name
				tag := t.Field(i).Tag.Get(Nut)
				err := checkStruct(fvc, fieldName, tag)
				if err != nil {
					return err
				}
			default:
				ft := t.Field(i)
				tv := ft.Tag.Get(Nut)
				cns := strings.Split(tv, Semicolon)
				err := check(fvc.Type(), cns)
				if err != nil {
					return err
				}
			}
		case reflect.Slice, reflect.Array:
			ft := t.Field(i)
			tv := ft.Tag.Get(Nut)
			cns := strings.Split(tv, Semicolon)
			err := checkSlice(t.Field(i).Name, cns)
			if err != nil {
				return err
			}
		default:
			ft := t.Field(i)
			tv := ft.Tag.Get(Nut)
			cns := strings.Split(tv, Semicolon)
			err := check(ft.Type, cns)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func check(field reflect.Type, tvs []string) error {
	fieldName := field.Name()
	switch field.Kind() {
	case reflect.Int8:
		return checkInt8(fieldName, tvs)
	case reflect.Uint8:
		return checkUint8(fieldName, tvs)
	case reflect.Int16:
		return checkInt16(fieldName, tvs)
	case reflect.Uint16:
		return checkUint16(fieldName, tvs)
	case reflect.Int32:
		return checkInt32(fieldName, tvs)
	case reflect.Uint32:
		return checkUint32(fieldName, tvs)
	case reflect.Int:
		return checkInt(fieldName, tvs)
	case reflect.Uint:
		return checkUint(fieldName, tvs)
	case reflect.Int64:
		return checkInt64(fieldName, tvs)
	case reflect.Uint64:
		return checkUint64(fieldName, tvs)
	case reflect.Float32:
		return checkFloat32(fieldName, tvs)
	case reflect.Float64:
		return checkFloat64(fieldName, tvs)
	case reflect.String:
		return checkString(fieldName, tvs)
	case reflect.Slice:
		return checkSlice(fieldName, tvs)
	case reflect.Array:
		return checkSlice(fieldName, tvs)
	case reflect.Bool:
		return checkBool(fieldName, tvs)

	}

	return nil
}

func checkInt8(fieldName string, tvs []string) error {
	for _, v := range tvs {
		if v == "" {
			continue
		}
		if v == Optional || v == Required {
			continue
		}
		cds := strings.Split(v, ":")
		if len(cds) != 2 {
			return condValNumErr(fieldName, len(cds), v)
		}
		if !ArrayContains(NumeralCondSet[:], cds[0]) {
			return condKeyErr(fieldName)
		}
		cds0 := cds[0]
		switch cds0 {
		case Eq, Neq, Gt, Gte, Lt, Lte:
			_, err := strconv.Atoi(cds[1])
			if err != nil {
				return condFormatErr(fieldName)
			}
		case Between:
			cvs := strings.Split(cds[1], ",")
			if len(cvs) != 2 {
				return condValNumErr(fieldName, len(cvs), cds0)
			}
			left, err := strconv.Atoi(cvs[0])
			if err != nil {
				return condFormatErr(fieldName)
			}
			right, err := strconv.Atoi(cvs[1])
			if err != nil {
				return condFormatErr(fieldName)
			}
			if left >= right {
				return condValLogicErr(fieldName, cds0)
			}
		case Size:
			err := checkSize(fieldName, cds[1])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func checkUint8(fieldName string, tvs []string) error {
	for _, v := range tvs {
		if v == "" {
			continue
		}
		if v == Optional || v == Required {
			continue
		}
		cds := strings.Split(v, ":")
		if len(cds) != 2 {
			return condValNumErr(fieldName, len(cds), v)
		}
		if !ArrayContains(NumeralCondSet[:], cds[0]) {
			return condKeyErr(fieldName)
		}
		cds0 := cds[0]
		switch cds0 {
		case Eq, Neq, Gt, Gte, Lt, Lte:
			_, err := strconv.Atoi(cds[1])
			if err != nil {
				return condFormatErr(fieldName)
			}
		case Between:
			cvs := strings.Split(cds[1], ",")
			if len(cvs) != 2 {
				return condValNumErr(fieldName, len(cvs), cds0)
			}
			left, err := strconv.Atoi(cvs[0])
			if err != nil {
				return condFormatErr(fieldName)
			}
			right, err := strconv.Atoi(cvs[1])
			if err != nil {
				return condFormatErr(fieldName)
			}
			if left >= right {
				return condValLogicErr(fieldName, cds0)
			}
		case Size:
			err := checkSize(fieldName, cds[1])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func checkInt16(fieldName string, tvs []string) error {
	for _, v := range tvs {
		if v == "" {
			continue
		}
		if v == Optional || v == Required {
			continue
		}
		cds := strings.Split(v, ":")
		if len(cds) != 2 {
			return condValNumErr(fieldName, len(cds), v)
		}
		if !ArrayContains(NumeralCondSet[:], cds[0]) {
			return condKeyErr(fieldName)
		}
		cds0 := cds[0]
		switch cds0 {
		case Eq, Neq, Gt, Gte, Lt, Lte:
			_, err := strconv.Atoi(cds[1])
			if err != nil {
				return condFormatErr(fieldName)
			}
		case Between:
			cvs := strings.Split(cds[1], ",")
			if len(cvs) != 2 {
				return condValNumErr(fieldName, len(cvs), cds0)
			}
			left, err := strconv.Atoi(cvs[0])
			if err != nil {
				return condFormatErr(fieldName)
			}
			right, err := strconv.Atoi(cvs[1])
			if err != nil {
				return condFormatErr(fieldName)
			}
			if left >= right {
				return condValLogicErr(fieldName, cds0)
			}
		case Size:
			err := checkSize(fieldName, cds[1])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func checkUint16(fieldName string, tvs []string) error {
	for _, v := range tvs {
		if v == "" {
			continue
		}
		if v == Optional || v == Required {
			continue
		}
		cds := strings.Split(v, ":")
		if len(cds) != 2 {
			return condValNumErr(fieldName, len(cds), v)
		}
		if !ArrayContains(NumeralCondSet[:], cds[0]) {
			return condKeyErr(fieldName)
		}
		cds0 := cds[0]
		switch cds0 {
		case Eq, Neq, Gt, Gte, Lt, Lte:
			_, err := strconv.Atoi(cds[1])
			if err != nil {
				return condFormatErr(fieldName)
			}
		case Between:
			cvs := strings.Split(cds[1], ",")
			if len(cvs) != 2 {
				return condValNumErr(fieldName, len(cvs), cds0)
			}
			left, err := strconv.Atoi(cvs[0])
			if err != nil {
				return condFormatErr(fieldName)
			}
			right, err := strconv.Atoi(cvs[1])
			if err != nil {
				return condFormatErr(fieldName)
			}
			if left >= right {
				return condValLogicErr(fieldName, cds0)
			}
		case Size:
			err := checkSize(fieldName, cds[1])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func checkInt32(fieldName string, tvs []string) error {
	for _, v := range tvs {
		if v == "" {
			continue
		}
		if v == Optional || v == Required {
			continue
		}
		cds := strings.Split(v, ":")
		if len(cds) != 2 {
			return condValNumErr(fieldName, len(cds), v)
		}
		if !ArrayContains(NumeralCondSet[:], cds[0]) {
			return condKeyErr(fieldName)
		}
		cds0 := cds[0]
		switch cds0 {
		case Eq, Neq, Gt, Gte, Lt, Lte:
			_, err := strconv.Atoi(cds[1])
			if err != nil {
				return condFormatErr(fieldName)
			}
		case Between:
			cvs := strings.Split(cds[1], ",")
			if len(cvs) != 2 {
				return condValNumErr(fieldName, len(cvs), cds0)
			}
			left, err := strconv.Atoi(cvs[0])
			if err != nil {
				return condFormatErr(fieldName)
			}
			right, err := strconv.Atoi(cvs[1])
			if err != nil {
				return condFormatErr(fieldName)
			}
			if left >= right {
				return condValLogicErr(fieldName, cds0)
			}
		case Size:
			err := checkSize(fieldName, cds[1])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func checkUint32(fieldName string, tvs []string) error {
	for _, v := range tvs {
		if v == "" {
			continue
		}
		if v == Optional || v == Required {
			continue
		}
		cds := strings.Split(v, ":")
		if len(cds) != 2 {
			return condValNumErr(fieldName, len(cds), v)
		}
		if !ArrayContains(NumeralCondSet[:], cds[0]) {
			return condKeyErr(fieldName)
		}
		cds0 := cds[0]
		switch cds0 {
		case Eq, Neq, Gt, Gte, Lt, Lte:
			_, err := strconv.Atoi(cds[1])
			if err != nil {
				return condFormatErr(fieldName)
			}
		case Between:
			cvs := strings.Split(cds[1], ",")
			if len(cvs) != 2 {
				return condValNumErr(fieldName, len(cvs), cds0)
			}
			left, err := strconv.Atoi(cvs[0])
			if err != nil {
				return condFormatErr(fieldName)
			}
			right, err := strconv.Atoi(cvs[1])
			if err != nil {
				return condFormatErr(fieldName)
			}
			if left >= right {
				return condValLogicErr(fieldName, cds0)
			}
		case Size:
			err := checkSize(fieldName, cds[1])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func checkInt(fieldName string, tvs []string) error {
	for _, v := range tvs {
		if v == "" {
			continue
		}
		if v == Optional || v == Required {
			continue
		}
		cds := strings.Split(v, ":")
		if len(cds) != 2 {
			return condValNumErr(fieldName, len(cds), v)
		}
		if !ArrayContains(NumeralCondSet[:], cds[0]) {
			return condKeyErr(fieldName)
		}
		cds0 := cds[0]
		switch cds0 {
		case Eq, Neq, Gt, Gte, Lt, Lte:
			_, err := strconv.Atoi(cds[1])
			if err != nil {
				return condFormatErr(fieldName)
			}
		case Between:
			cvs := strings.Split(cds[1], ",")
			if len(cvs) != 2 {
				return condValNumErr(fieldName, len(cvs), cds0)
			}
			left, err := strconv.Atoi(cvs[0])
			if err != nil {
				return condFormatErr(fieldName)
			}
			right, err := strconv.Atoi(cvs[1])
			if err != nil {
				return condFormatErr(fieldName)
			}
			if left >= right {
				return condValLogicErr(fieldName, cds0)
			}
		case Size:
			err := checkSize(fieldName, cds[1])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func checkUint(fieldName string, tvs []string) error {
	for _, v := range tvs {
		if v == "" {
			continue
		}
		if v == Optional || v == Required {
			continue
		}
		cds := strings.Split(v, ":")
		if len(cds) != 2 {
			return condValNumErr(fieldName, len(cds), v)
		}
		if !ArrayContains(NumeralCondSet[:], cds[0]) {
			return condKeyErr(fieldName)
		}
		cds0 := cds[0]
		switch cds0 {
		case Eq, Neq, Gt, Gte, Lt, Lte:
			_, err := strconv.ParseUint(cds[1], 10, 64)
			if err != nil {
				return condFormatErr(fieldName)
			}
		case Between:
			cvs := strings.Split(cds[1], ",")
			if len(cvs) != 2 {
				return condValNumErr(fieldName, len(cvs), cds0)
			}
			left, err := strconv.ParseUint(cvs[0], 10, 64)
			if err != nil {
				return condFormatErr(fieldName)
			}
			right, err := strconv.ParseUint(cvs[1], 10, 64)
			if err != nil {
				return condFormatErr(fieldName)
			}
			if left >= right {
				return condValLogicErr(fieldName, cds0)
			}
		case Size:
			err := checkSize(fieldName, cds[1])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func checkInt64(fieldName string, tvs []string) error {
	for _, v := range tvs {
		if v == "" {
			continue
		}
		if v == Optional || v == Required {
			continue
		}
		cds := strings.Split(v, ":")
		if len(cds) != 2 {
			return condValNumErr(fieldName, len(cds), v)
		}
		if !ArrayContains(NumeralCondSet[:], cds[0]) {
			return condKeyErr(fieldName)
		}
		cds0 := cds[0]
		switch cds0 {
		case Eq, Neq, Gt, Gte, Lt, Lte:
			_, err := strconv.ParseInt(cds[1], 10, 64)
			if err != nil {
				return condFormatErr(fieldName)
			}
		case Between:
			cvs := strings.Split(cds[1], ",")
			if len(cvs) != 2 {
				return condValNumErr(fieldName, len(cvs), cds0)
			}
			left, err := strconv.ParseInt(cvs[0], 10, 64)
			if err != nil {
				return condFormatErr(fieldName)
			}
			right, err := strconv.ParseInt(cvs[1], 10, 64)
			if err != nil {
				return condFormatErr(fieldName)
			}
			if left >= right {
				return condValLogicErr(fieldName, cds0)
			}
		case Size:
			err := checkSize(fieldName, cds[1])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func checkUint64(fieldName string, tvs []string) error {
	for _, v := range tvs {
		if v == "" {
			continue
		}
		if v == Optional || v == Required {
			continue
		}
		cds := strings.Split(v, ":")
		if len(cds) != 2 {
			return condValNumErr(fieldName, len(cds), v)
		}
		if !ArrayContains(NumeralCondSet[:], cds[0]) {
			return condKeyErr(fieldName)
		}
		cds0 := cds[0]
		switch cds0 {
		case Eq, Neq, Gt, Gte, Lt, Lte:
			_, err := strconv.ParseUint(cds[1], 10, 64)
			if err != nil {
				return condFormatErr(fieldName)
			}
		case Between:
			cvs := strings.Split(cds[1], ",")
			if len(cvs) != 2 {
				return condValNumErr(fieldName, len(cvs), cds0)
			}
			left, err := strconv.ParseUint(cvs[0], 10, 64)
			if err != nil {
				return condFormatErr(fieldName)
			}
			right, err := strconv.ParseUint(cvs[1], 10, 64)
			if err != nil {
				return condFormatErr(fieldName)
			}
			if left >= right {
				return condValLogicErr(fieldName, cds0)
			}
		case Size:
			err := checkSize(fieldName, cds[1])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func checkFloat32(fieldName string, tvs []string) error {
	for _, v := range tvs {
		if v == "" {
			continue
		}
		if v == Optional || v == Required {
			continue
		}
		cds := strings.Split(v, ":")
		if len(cds) != 2 {
			return condValNumErr(fieldName, len(cds), v)
		}

		if !ArrayContains(NumeralCondSet[:], cds[0]) {
			return condKeyErr(fieldName)
		}
		cds0 := cds[0]
		switch cds0 {
		case Eq, Neq, Gt, Gte, Lt, Lte:
			_, err := strconv.ParseFloat(cds[1], 32)
			if err != nil {
				return condFormatErr(fieldName)
			}
		case Between:
			cvs := strings.Split(cds[1], ",")
			if len(cvs) != 2 {
				return condValNumErr(fieldName, len(cvs), cds0)
			}
			left, err := strconv.ParseFloat(cvs[0], 32)
			if err != nil {
				return condFormatErr(fieldName)
			}
			right, err := strconv.ParseFloat(cvs[1], 32)
			if err != nil {
				return condFormatErr(fieldName)
			}
			if left >= right {
				return condValLogicErr(fieldName, cds0)
			}
		case Size:
			err := checkSize(fieldName, cds[1])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func checkFloat64(fieldName string, tvs []string) error {
	for _, v := range tvs {
		if v == "" {
			continue
		}
		if v == Optional || v == Required {
			continue
		}
		cds := strings.Split(v, ":")
		if len(cds) != 2 {
			return condValNumErr(fieldName, len(cds), v)
		}

		if !ArrayContains(NumeralCondSet[:], cds[0]) {
			return condKeyErr(fieldName)
		}

		cds0 := cds[0]
		switch cds0 {
		case Eq, Neq, Gt, Gte, Lt, Lte:
			_, err := strconv.ParseFloat(cds[1], 64)
			if err != nil {
				return condFormatErr(fieldName)
			}
		case Between:
			cvs := strings.Split(cds[1], ",")
			if len(cvs) != 2 {
				return condValNumErr(fieldName, len(cvs), cds0)
			}
			left, err := strconv.ParseFloat(cvs[0], 64)
			if err != nil {
				return condFormatErr(fieldName)
			}
			right, err := strconv.ParseFloat(cvs[1], 64)
			if err != nil {
				return condFormatErr(fieldName)
			}
			if left >= right {
				return condValLogicErr(fieldName, cds0)
			}
		case Size:
			err := checkSize(fieldName, cds[1])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func checkString(fieldName string, tvs []string) error {
	for _, v := range tvs {
		if v == "" {
			continue
		}
		if v == Optional || v == Required {
			continue
		}
		cds := strings.Split(v, ":")
		if len(cds) != 2 {
			return condValNumErr(fieldName, len(cds), v)
		}
		if !ArrayContains(StringCondSet[:], cds[0]) {
			return condKeyErr(fieldName)
		}
		switch cds[0] {
		case Size:
			err := checkSize(fieldName, cds[1])
			if err != nil {
				return err
			}
		case In:
			if cds[1] == "" {
				return condValLogicErr(fieldName, cds[0])
			}
		case Regexp:
			if cds[1] == "" {
				return condValLogicErr(fieldName, cds[0])
			}
		}
	}

	return nil
}

func checkSlice(fieldName string, tvs []string) error {
	for _, v := range tvs {
		if v == "" {
			continue
		}
		if v == Optional || v == Required {
			continue
		}
		cds := strings.Split(v, ":")
		if len(cds) != 2 {
			return condValNumErr(fieldName, len(cds), v)
		}
		if !ArrayContains(ArrayCondSet[:], cds[0]) {
			return condKeyErr(fieldName)
		}
		switch cds[0] {
		case Size:
			err := checkSize(fieldName, cds[1])
			if err != nil {
				return err
			}
		case Contains, Excluded:
			var elemType string
			for _, m := range tvs {
				rule := strings.Split(m, ":")
				if rule[0] == Type {
					elemType = rule[1]
				}
			}
			if elemType == "" {
				return condNeedType(fieldName)
			}
			if !ArrayContains(Types[:], elemType) {
				return condNeedType(fieldName)
			}
		}
	}

	return nil
}

func checkBool(fieldName string, tvs []string) error {
	for _, v := range tvs {
		if v == "" {
			continue
		}
		cds := strings.Split(v, ":")
		if len(cds) != 2 {
			return condValNumErr(fieldName, len(cds), v)
		}
		if !ArrayContains(BoolCondSet[:], cds[0]) {
			return condKeyErr(fieldName)
		}
	}

	return nil
}

func checkStruct(v reflect.Value, fieldName string, tag string) error {
	if !v.CanInterface() {
		return errors.New("exists on an unsafe type")
	}

	if tag == "" {
		return nil
	}

	tvs := strings.Split(tag, Semicolon)
	if len(tvs) != 1 {
		return condKeyErr(fieldName)
	}

	if !ArrayContains(StructCondSet[:], tvs[0]) {
		return condKeyErr(fieldName)
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		ft := t.Field(i)
		tv := ft.Tag.Get(Nut)
		cns := strings.Split(tv, ";")
		err := check(ft.Type, cns)
		if err != nil {
			return err
		}

	}

	return nil
}

func checkRule(v reflect.Value) error {
	cc1 := make(map[string]int8)
	cc2 := make(map[string]int8)
	ft := v.Type()
	for i := 0; i < v.NumField(); i++ {
		tfd := ft.Field(i)
		definedRules := getDefinedRules(tfd.Tag.Get(Nut))
		_, ok1 := definedRules[Optional]
		_, ok2 := definedRules[Required]
		if ok1 && ok2 {
			return fmt.Errorf("optional and required can only have one")
		}
		vfd := v.Field(i)
		switch vfd.Kind() {
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int, reflect.Int64,
			reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint, reflect.Uint64,
			reflect.Float32, reflect.Float64:
			for k := range definedRules {
				if !ArrayContains(NumeralCondSet[:], k) {
					return condKeyErr(tfd.Name)
				}
				if ArrayContains([]string{In, Type, Size, Optional, Required}, k) {
					continue
				}
				cc1[tfd.Name] += 1
				if cc1[tfd.Name] > 1 {
					return condRuleErr(tfd.Name)
				}
			}
		case reflect.String:
			for k := range definedRules {
				if !ArrayContains(StringCondSet[:], k) {
					return condKeyErr(tfd.Name)
				}
				if ArrayContains([]string{In, Type, Regexp, Size, Optional, Required}, k) {
					continue
				}
				cc1[tfd.Name] += 1
				if cc1[tfd.Name] > 1 {
					return condRuleErr(tfd.Name)
				}
			}
		case reflect.Slice, reflect.Array:
			for k := range definedRules {
				if !ArrayContains(ArrayCondSet[:], k) {
					return condKeyErr(tfd.Name)
				}
				if ArrayContains([]string{Type, Size, Optional, Required}, k) {
					continue
				}
				cc2[tfd.Name] += 1
				if cc2[tfd.Name] > 1 {
					return condRuleErr(tfd.Name)
				}
				if k == Contains || k == Excluded {
					v, ok := definedRules[Type]
					if !ok {
						return fmt.Errorf(`%s condition needs to specify "type"`, k)
					}
					if !ArrayContains(Types[:], v) {
						return fmt.Errorf(`the "type" specified in the %s field is wrong`, tfd.Name)
					}
				}
			}
		}
	}

	return nil
}

func checkSize(fieldName, cValue string) error {
	sv := strings.Split(cValue, ",")
	if len(sv) == 2 {
		min, err := strconv.ParseUint(sv[0], 10, 64)
		if err != nil {
			return condFormatErr(fieldName)
		}
		max, err := strconv.ParseUint(sv[1], 10, 64)
		if err != nil {
			return condFormatErr(fieldName)
		}
		if min >= max {
			return condValLogicErr(fieldName, Size)
		}
	} else if len(sv) == 1 {
		_, err := strconv.ParseUint(sv[0], 10, 64)
		if err != nil {
			return condFormatErr(fieldName)
		}
	} else {
		return condValNumErr(fieldName, len(sv), Size)
	}

	return nil
}

//condition error function

func condFormatErr(fieldName string) error {
	return fmt.Errorf("the conditional value of the %q field is in the wrong format", fieldName)
}

func condKeyErr(fieldName string) error {
	return fmt.Errorf("the condition type of the %q field is wrong", fieldName)
}

func condRuleErr(fieldName string) error {
	return fmt.Errorf("the verification rule of field %s is wrong, there can only be one comparison condition", fieldName)
}

func condValNumErr(fieldName string, n int, kind string) error {
	return fmt.Errorf("the condition of the %q field is %s, but %d values are set", fieldName, kind, n)
}

func condValLogicErr(fieldName string, kind string) error {
	return fmt.Errorf("the condition of field %q is %s, but the condition value logic is wrong", fieldName, kind)
}

func condNeedType(fieldName string) error {
	return fmt.Errorf(`the conditional value of the %q field need tag "type"`, fieldName)
}
