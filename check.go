package nut

import (
	"errors"
	"fmt"
	"math"
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
			err := checkStruct(fv)
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
				err := checkStruct(fvc)
				if err != nil {
					return err
				}
			default:
				ft := t.Field(i)
				tv := ft.Tag.Get(Nut)
				cns := strings.Split(tv, ";")
				err := check(fvc.Type(), cns)
				if err != nil {
					return err
				}
			}
		case reflect.Slice, reflect.Array:
			ft := t.Field(i)
			tv := ft.Tag.Get(Nut)
			cns := strings.Split(tv, ";")
			err := checkSlice(t.Field(i).Name, cns)
			if err != nil {
				return err
			}
		default:
			ft := t.Field(i)
			tv := ft.Tag.Get(Nut)
			cns := strings.Split(tv, ";")
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
			tv, err := strconv.Atoi(cds[1])
			if err != nil {
				return condFormatErr(fieldName)
			}
			if tv > math.MaxInt8 || tv < math.MinInt8 {
				return overflowErr(fieldName)
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
			if left > math.MaxInt8 || left < math.MinInt8 {
				return overflowErr(fieldName)
			}
			right, err := strconv.Atoi(cvs[1])
			if err != nil {
				return condFormatErr(fieldName)
			}
			if right > math.MaxInt8 || right < math.MinInt8 {
				return overflowErr(fieldName)
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
			tv, err := strconv.Atoi(cds[1])
			if err != nil {
				return condFormatErr(fieldName)
			}
			if tv > math.MaxUint8 || tv < Zero {
				return overflowErr(fieldName)
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
			if left > math.MaxUint8 || left < Zero {
				return overflowErr(fieldName)
			}
			right, err := strconv.Atoi(cvs[1])
			if err != nil {
				return condFormatErr(fieldName)
			}
			if right > math.MaxUint8 || right < Zero {
				return overflowErr(fieldName)
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
			tv, err := strconv.Atoi(cds[1])
			if err != nil {
				return condFormatErr(fieldName)
			}
			if tv > math.MaxInt16 || tv < math.MinInt16 {
				return overflowErr(fieldName)
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
			if left > math.MaxInt16 || left < math.MinInt16 {
				return overflowErr(fieldName)
			}
			right, err := strconv.Atoi(cvs[1])
			if err != nil {
				return condFormatErr(fieldName)
			}
			if right > math.MaxInt16 || right < math.MinInt16 {
				return overflowErr(fieldName)
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
			tv, err := strconv.Atoi(cds[1])
			if err != nil {
				return condFormatErr(fieldName)
			}
			if tv > math.MaxUint16 || tv < Zero {
				return overflowErr(fieldName)
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
			if left > math.MaxUint16 || left < Zero {
				return overflowErr(fieldName)
			}
			right, err := strconv.Atoi(cvs[1])
			if err != nil {
				return condFormatErr(fieldName)
			}
			if right > math.MaxUint16 || right < Zero {
				return overflowErr(fieldName)
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
			tv, err := strconv.Atoi(cds[1])
			if err != nil {
				return condFormatErr(fieldName)
			}
			if tv > math.MaxInt32 || tv < math.MinInt32 {
				return overflowErr(fieldName)
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
			if left > math.MaxInt32 || left < math.MinInt32 {
				return overflowErr(fieldName)
			}
			right, err := strconv.Atoi(cvs[1])
			if err != nil {
				return condFormatErr(fieldName)
			}
			if right > math.MaxInt32 || right < math.MinInt32 {
				return overflowErr(fieldName)
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
			tv, err := strconv.Atoi(cds[1])
			if err != nil {
				return condFormatErr(fieldName)
			}
			if tv > math.MaxUint32 || tv < Zero {
				return overflowErr(fieldName)
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
			if left > math.MaxUint32 || left < Zero {
				return overflowErr(fieldName)
			}
			right, err := strconv.Atoi(cvs[1])
			if err != nil {
				return condFormatErr(fieldName)
			}
			if right > math.MaxUint32 || right < Zero {
				return overflowErr(fieldName)
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
			tv, err := strconv.Atoi(cds[1])
			if err != nil {
				return condFormatErr(fieldName)
			}
			if tv > math.MaxInt || tv < math.MinInt {
				return overflowErr(fieldName)
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
			if left > math.MaxInt || left < math.MinInt {
				return overflowErr(fieldName)
			}
			right, err := strconv.Atoi(cvs[1])
			if err != nil {
				return condFormatErr(fieldName)
			}
			if right > math.MaxInt || right < math.MinInt {
				return overflowErr(fieldName)
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
			tv, err := strconv.ParseUint(cds[1], 10, 64)
			if err != nil {
				return condFormatErr(fieldName)
			}
			if tv > math.MaxUint || tv < Zero {
				return overflowErr(fieldName)
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
			if left > math.MaxUint || left < Zero {
				return overflowErr(fieldName)
			}
			right, err := strconv.ParseUint(cvs[1], 10, 64)
			if err != nil {
				return condFormatErr(fieldName)
			}
			if right > math.MaxUint || right < Zero {
				return overflowErr(fieldName)
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
			tv, err := strconv.ParseInt(cds[1], 10, 64)
			if err != nil {
				return condFormatErr(fieldName)
			}
			if tv > math.MaxInt64 || tv < math.MinInt64 {
				return overflowErr(fieldName)
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
			if left > math.MaxInt64 || left < math.MinInt64 {
				return overflowErr(fieldName)
			}
			right, err := strconv.ParseInt(cvs[1], 10, 64)
			if err != nil {
				return condFormatErr(fieldName)
			}
			if right > math.MaxInt64 || right < math.MinInt64 {
				return overflowErr(fieldName)
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
			tv, err := strconv.ParseUint(cds[1], 10, 64)
			if err != nil {
				return condFormatErr(fieldName)
			}
			if tv > math.MaxUint || tv < Zero {
				return overflowErr(fieldName)
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
			if left > math.MaxUint || left < Zero {
				return overflowErr(fieldName)
			}
			right, err := strconv.ParseUint(cvs[1], 10, 64)
			if err != nil {
				return condFormatErr(fieldName)
			}
			if right > math.MaxUint || right < Zero {
				return overflowErr(fieldName)
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
			tv, err := strconv.ParseFloat(cds[1], 32)
			if err != nil {
				return condFormatErr(fieldName)
			}
			if tv > math.MaxFloat32 || tv < math.SmallestNonzeroFloat32 {
				return overflowErr(fieldName)
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
			if left > math.MaxFloat32 || left < math.SmallestNonzeroFloat32 {
				return overflowErr(fieldName)
			}
			right, err := strconv.ParseFloat(cvs[1], 32)
			if err != nil {
				return condFormatErr(fieldName)
			}
			if right > math.MaxFloat32 || right < math.SmallestNonzeroFloat32 {
				return overflowErr(fieldName)
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
			tv, err := strconv.ParseFloat(cds[1], 64)
			if err != nil {
				return condFormatErr(fieldName)
			}
			if tv > math.MaxFloat64 || tv < math.SmallestNonzeroFloat64 {
				return overflowErr(fieldName)
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
			if left > math.MaxFloat64 || left < math.SmallestNonzeroFloat64 {
				return overflowErr(fieldName)
			}
			right, err := strconv.ParseFloat(cvs[1], 64)
			if err != nil {
				return condFormatErr(fieldName)
			}
			if right > math.MaxFloat64 || right < math.SmallestNonzeroFloat64 {
				return overflowErr(fieldName)
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
		cds := strings.Split(v, ":")
		if len(cds) != 2 {
			return condValNumErr(fieldName, len(cds), v)
		}
		if !ArrayContains(ArrayCondSet[:], cds[0]) {
			return condKeyErr(fieldName)
		}
		if cds[0] == Size {
			err := checkSize(fieldName, cds[1])
			if err != nil {
				return err
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

func checkStruct(v reflect.Value) error {
	if !v.CanInterface() {
		return errors.New("exists on an unsafe type")
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
		vfd := v.Field(i)
		tfd := ft.Field(i)
		switch vfd.Kind() {
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int, reflect.Int64,
			reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint, reflect.Uint64,
			reflect.Float32, reflect.Float64:
			rules := strings.Split(tfd.Tag.Get(Nut), ";")
			for _, v := range rules {
				rule := strings.Split(v, ":")
				if !ArrayContains(NumeralCondSet[:], rule[0]) {
					return condKeyErr(tfd.Name)
				}
				if ArrayContains([]string{In, Type, Size}, rule[0]) {
					continue
				}
				cc1[tfd.Name] += 1
				if cc1[tfd.Name] > 1 {
					return condRuleErr(tfd.Name)
				}
			}
		case reflect.String:
			rules := strings.Split(tfd.Tag.Get(Nut), ";")
			for _, v := range rules {
				rule := strings.Split(v, ":")
				if !ArrayContains(StringCondSet[:], rule[0]) {
					return condKeyErr(tfd.Name)
				}
				if ArrayContains([]string{In, Type, Regexp, Size}, rule[0]) {
					continue
				}
				cc1[tfd.Name] += 1
				if cc1[tfd.Name] > 1 {
					return condRuleErr(tfd.Name)
				}
			}
		case reflect.Slice, reflect.Array:
			rules := strings.Split(tfd.Tag.Get(Nut), ";")
			for _, v := range rules {
				rule := strings.Split(v, ":")
				if !ArrayContains(ArrayCondSet[:], rule[0]) {
					return condKeyErr(tfd.Name)
				}
				if ArrayContains([]string{Type, Size}, rule[0]) {
					continue
				}
				cc2[tfd.Name] += 1
				if cc2[tfd.Name] > 1 {
					return condRuleErr(tfd.Name)
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

func overflowErr(fieldName string) error {
	return fmt.Errorf("the conditional value of the %q field overflowed", fieldName)
}
