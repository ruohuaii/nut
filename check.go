package nut

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

func check(field reflect.StructField, tvs []string) error {
	fieldName := field.Name
	switch field.Type.Kind() {
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
	case reflect.Array:
	}
	return nil
}

func checkInt8(fieldName string, tvs []string) error {
	for _, v := range tvs {
		cds := strings.Split(v, ":")
		if len(cds) != 2 {
			return condValNumErr(fieldName, len(cds), v)
		}

		if !stringArray(conditions[:7]).Contains(cds[0]) {
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
		}
	}

	return nil
}

func checkUint8(fieldName string, tvs []string) error {
	for _, v := range tvs {
		cds := strings.Split(v, ":")
		if len(cds) != 2 {
			return condValNumErr(fieldName, len(cds), v)
		}

		if !stringArray(conditions[:7]).Contains(cds[0]) {
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
		}
	}

	return nil
}

func checkInt16(fieldName string, tvs []string) error {
	for _, v := range tvs {
		cds := strings.Split(v, ":")
		if len(cds) != 2 {
			return condValNumErr(fieldName, len(cds), v)
		}

		if !stringArray(conditions[:7]).Contains(cds[0]) {
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
		}
	}

	return nil
}

func checkUint16(fieldName string, tvs []string) error {
	for _, v := range tvs {
		cds := strings.Split(v, ":")
		if len(cds) != 2 {
			return condValNumErr(fieldName, len(cds), v)
		}

		if !stringArray(conditions[:7]).Contains(cds[0]) {
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
		}
	}

	return nil
}

func checkInt32(fieldName string, tvs []string) error {
	for _, v := range tvs {
		cds := strings.Split(v, ":")
		if len(cds) != 2 {
			return condValNumErr(fieldName, len(cds), v)
		}

		if !stringArray(conditions[:7]).Contains(cds[0]) {
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
		}
	}

	return nil
}

func checkUint32(fieldName string, tvs []string) error {
	for _, v := range tvs {
		cds := strings.Split(v, ":")
		if len(cds) != 2 {
			return condValNumErr(fieldName, len(cds), v)
		}

		if !stringArray(conditions[:7]).Contains(cds[0]) {
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
		}
	}

	return nil
}

func checkInt(fieldName string, tvs []string) error {
	for _, v := range tvs {
		cds := strings.Split(v, ":")
		if len(cds) != 2 {
			return condValNumErr(fieldName, len(cds), v)
		}

		if !stringArray(conditions[:7]).Contains(cds[0]) {
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
		}
	}

	return nil
}

func checkUint(fieldName string, tvs []string) error {
	for _, v := range tvs {
		cds := strings.Split(v, ":")
		if len(cds) != 2 {
			return condValNumErr(fieldName, len(cds), v)
		}

		if !stringArray(conditions[:7]).Contains(cds[0]) {
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
		}
	}

	return nil
}

func checkInt64(fieldName string, tvs []string) error {
	for _, v := range tvs {
		cds := strings.Split(v, ":")
		if len(cds) != 2 {
			return condValNumErr(fieldName, len(cds), v)
		}

		if !stringArray(conditions[:7]).Contains(cds[0]) {
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
		}
	}

	return nil
}

func checkUint64(fieldName string, tvs []string) error {
	for _, v := range tvs {
		cds := strings.Split(v, ":")
		if len(cds) != 2 {
			return condValNumErr(fieldName, len(cds), v)
		}

		if !stringArray(conditions[:7]).Contains(cds[0]) {
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
		}
	}

	return nil
}

func checkFloat32(fieldName string, tvs []string) error {
	for _, v := range tvs {
		cds := strings.Split(v, ":")
		if len(cds) != 2 {
			return condValNumErr(fieldName, len(cds), v)
		}

		if !stringArray(conditions[:7]).Contains(cds[0]) {
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
		}
	}

	return nil
}

func checkFloat64(fieldName string, tvs []string) error {
	for _, v := range tvs {
		cds := strings.Split(v, ":")
		if len(cds) != 2 {
			return condValNumErr(fieldName, len(cds), v)
		}

		if !stringArray(conditions[:7]).Contains(cds[0]) {
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
		}
	}

	return nil
}

func checkString(fieldName string, tvs []string) error {
	for _, v := range tvs {
		cds := strings.Split(v, ":")
		if len(cds) != 2 {
			return condValNumErr(fieldName, len(cds), v)
		}

		if !stringArray(conditions[:7]).Contains(cds[0]) {
			return condKeyErr(fieldName)
		}
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

func condValNumErr(fieldName string, n int, kind string) error {
	return fmt.Errorf("the condition of the %q field is %s, but %d values are set", fieldName, kind, n)
}

func condValLogicErr(fieldName string, kind string) error {
	return fmt.Errorf("the condition of field %q is %s, but the condition value logic is wrong", fieldName, kind)
}

func overflowErr(fieldName string) error {
	return fmt.Errorf("the conditional value of the %q field overflowed", fieldName)
}
