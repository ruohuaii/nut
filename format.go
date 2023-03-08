package nut

import (
	"fmt"
	"strings"
)

func ThrowCondEq(shortName, fieldName string, cValue any, isString bool) string {
	var format string
	if !isString {
		format = `if %s.%s == %v {
     	return fmt.Errorf("the value of %s field should not be equal to %v")
	}`
	} else {
		format = `if %s.%s == "%v" {
     	return fmt.Errorf("the value of %s field should not be equal to %v")
	}`
	}

	return fmt.Sprintf(format, shortName, fieldName, cValue, fieldName, cValue)
}

func ThrowCondNeq(shortName, fieldName string, cValue any, isString bool) string {
	var format string
	if !isString {
		format = `if %s.%s != %v {
     	return fmt.Errorf("the value of %s field should be equal to %v")
	}`
	} else {
		format = `if %s.%s != "%v" {
     	return fmt.Errorf("the value of %s field should be equal to %v")
	}`
	}

	return fmt.Sprintf(format, shortName, fieldName, cValue, fieldName, cValue)
}

func ThrowCondLt(shortName, fieldName string, cValue any) string {
	format := `if %s.%s < %v {
		return fmt.Errorf("the value of %s field should be greater than or equal to %v")
	}`
	return fmt.Sprintf(format, shortName, fieldName, cValue, fieldName, cValue)
}

func ThrowCondLte(shortName, fieldName string, cValue any) string {
	format := `if %s.%s <= %v {
		return fmt.Errorf("the value of %s field should be greater than %v")
	}`
	return fmt.Sprintf(format, shortName, fieldName, cValue, fieldName, cValue)
}

func ThrowCondGt(shortName, fieldName string, cValue any) string {
	format := `if %s.%s > %v {
		return fmt.Errorf("the value of %s field should be less than or equal to %v")
	}`
	return fmt.Sprintf(format, shortName, fieldName, cValue, fieldName, cValue)
}

func ThrowCondGte(shortName, fieldName string, cValue any) string {
	format := `if %s.%s >= %v {
		return fmt.Errorf("the value of %s field should be less than %v")
	}`
	return fmt.Sprintf(format, shortName, fieldName, cValue, fieldName, cValue)
}

func ThrowCondBetween(shortName, fieldName string, left, right any) string {
	format := `if %s.%s < %v || %s.%s > %v {
     	return fmt.Errorf("the value of %s field should be between %v and %v")
	}`
	return fmt.Sprintf(format, shortName, fieldName, left, shortName, fieldName, right,
		fieldName, left, right)
}

func ThrowCondRegexp(shortName, fieldName string, matchRule string) string {
	varName := fmt.Sprintf("%sRegexp", strings.ToLower(fieldName))
	format := `%s,err := regexp.Compile(%q)
	if err != nil {
		return err
	}
	if !%s.MatchString(%s.%s) {
		return fmt.Errorf("the value of the %s field does not conform to the regular rules")
	}`
	return fmt.Sprintf(format, varName, matchRule, varName, shortName, fieldName, fieldName)
}

func ThrowCondSize(shortName, fieldName string, cvs []string, fieldType string) string {
	var format string
	if len(cvs) == 2 {
		switch fieldType {
		case "string", "array", "slice":
			format = `if len(%s.%s) < %v || len(%s.%s) > %v {
		return fmt.Errorf("the length of %s field value should be between %v and %v")
	}`
		case "int8", "int16", "int32", "int", "int64",
			"uint8", "uint16", "uint32", "uint", "uint64",
			"float32", "float64":
			format = `if len(fmt.Sprint(%s.%s)) < %v || len(fmt.Sprint(%s.%s)) > %v {
		return fmt.Errorf("the length of %s field value should be between %v and %v")
	}`
		}
		left, right := cvs[0], cvs[1]
		return fmt.Sprintf(format, shortName, fieldName, left, shortName, fieldName, right,
			fieldName, left, right)
	} else {
		switch fieldType {
		case "string", "array", "slice":
			format = `if len(%s.%s) != %s {
		return fmt.Errorf("the length of the %s field must be %s")
	}`
		case "int8", "int16", "int32", "int", "int64",
			"uint8", "uint16", "uint32", "uint", "uint64",
			"float32", "float64":
			format = `if len(fmt.Sprint(%s.%s)) != %s {
		return fmt.Errorf("the length of the %s field must be %s")
	}`
		}
		length := cvs[0]
		return fmt.Sprintf(format, shortName, fieldName, length, fieldName, length)
	}
}

func ThrowCondContains(shortName, fieldName string, cValue string, elemType string) string {
	elems := elemsBuilder(elemType, cValue)
	varName := fmt.Sprintf("%s%sContains", strings.ToLower(fieldName[:1]), fieldName[1:])
	format := `var %s= %s
	for i:=0;i<len(%s);i++{
		if ArrayContains(%s.%s,%s[i]) {
			return fmt.Errorf("the value of the %s field cannot contain %s")
		}
	}`

	return fmt.Sprintf(format, varName, elems, varName, shortName, fieldName, varName, fieldName, cValue)
}

func ThrowCondExcluded(shortName, fieldName string, cValue string, elemType string) string {
	elems := elemsBuilder(elemType, cValue)
	varName := fmt.Sprintf("%s%sExcluded", strings.ToLower(fieldName[:1]), fieldName[1:])
	format := `var %s = %s
	for i:=0;i<len(%s);i++{
		if !ArrayContains(%s.%s,%s[i]) {
			return fmt.Errorf("the value of the %s field must contain %s")
		}
	}`

	return fmt.Sprintf(format, varName, elems, varName, shortName, fieldName, varName, fieldName, cValue)
}

func ThrowCondIn(shortName, fieldName string, cValue string, elemType string) string {
	elems := elemsBuilder(elemType, cValue)
	varName := fmt.Sprintf("%s%sIn", strings.ToLower(fieldName[:1]), fieldName[1:])
	format := `var %s = %s
	if !ArrayContains(%s,%s.%s) {
		return fmt.Errorf("the value of the %s field should be one of %s")
	}`

	return fmt.Sprintf(format, varName, elems, varName, shortName, fieldName, fieldName, cValue)
}

func ThrowCondType(shortName, fieldName string, elemType string) string {
	var format string
	varName := fmt.Sprintf("%s%sVal", strings.ToLower(fieldName[:1]), fieldName[1:])
	switch elemType {
	case Int8:
		format = `%s,err := strconv.ParseInt(%s.%s,10,64)
	if err != nil {
		return fmt.Errorf("the value of the %s field is wrong")
	}
	if %s >math.MaxInt8||%s<math.MinInt8{
		return fmt.Errorf("the value of the %s field overflowed")
	}`
	case Int16:
		format = `%s,err := strconv.ParseInt(%s.%s,10,64)
	if err != nil {
		return fmt.Errorf("the value of the %s field is wrong")
	}
	if %s >math.MaxInt16||%s<math.MinInt16{
		return fmt.Errorf("the value of the %s field overflowed")
	}`
	case Int32:
		format = `%s,err := strconv.ParseInt(%s.%s,10,64)
	if err != nil {
		return fmt.Errorf("the value of the %s field is wrong")
	}
	if %s >math.MaxInt32||%s<math.MinInt32{
		return fmt.Errorf("the value of the %s field overflowed")
	}`
	case Int:
		format = `%s,err := strconv.ParseInt(%s.%s,10,64)
	if err != nil {
		return fmt.Errorf("the value of the %s field is wrong")
	}
	if %s >math.MaxInt||%s<math.MinInt{
		return fmt.Errorf("the value of the %s field overflowed")
	}`
	case Int64:
		format = `%s,err := strconv.ParseInt(%s.%s,10,64)
	if err != nil {
		return fmt.Errorf("the value of the %s field is wrong")
	}
	if %s >math.MaxInt64||%s<math.MinInt64{
		return fmt.Errorf("the value of the %s field overflowed")
	}`
	case Uint8:
		format = `%s,err := strconv.ParseUint(%s.%s,10,64)
	if err != nil {
		return fmt.Errorf("the value of the %s field is wrong")
	}
	if %s >math.MaxUint8||%s<0{
		return fmt.Errorf("the value of the %s field overflowed")
	}`
	case Uint16:
		format = `%s,err := strconv.ParseUint(%s.%s,10,64)
	if err != nil {
		return fmt.Errorf("the value of the %s field is wrong")
	}
	if %s >math.MaxUint16||%s<0{
		return fmt.Errorf("the value of the %s field overflowed")
	}`
	case Uint32:
		format = `%s,err := strconv.ParseUint(%s.%s,10,64)
	if err != nil {
		return fmt.Errorf("the value of the %s field is wrong")
	}
	if %s >math.MaxUint32||%s<0{
		return fmt.Errorf("the value of the %s field overflowed")
	}`
	case Uint64:
		format = `%s,err := strconv.ParseUint(%s.%s,10,64)
	if err != nil {
		return fmt.Errorf("the value of the %s field is wrong")
	}
	if %s >math.MaxUint64||%s<0{
		return fmt.Errorf("the value of the %s field overflowed")
	}`
	case Float32:
		format = `%s,err := strconv.ParseFloat(%s.%s,32)
	if err != nil {
		return fmt.Errorf("the value of the %s field is wrong")
	}
	if %s >math.MaxFloat32||%s<math.SmallestNonzeroFloat32{
		return fmt.Errorf("the value of the %s field overflowed")
	}`
	case Float64:
		format = `%s,err := strconv.ParseFloat(%s.%s,64)
	if err != nil {
		return fmt.Errorf("the value of the %s field is wrong")
	}
	if %s >math.MaxFloat64||%s<math.SmallestNonzeroFloat64{
		return fmt.Errorf("the value of the %s field overflowed")
	}`
	}

	return fmt.Sprintf(format, varName, shortName, fieldName, fieldName, varName, varName, fieldName)
}

func ThrowCondStruct(shortName, fieldName, structName string, isOptional bool, isPtr bool) string {
	var format string
	varName := fmt.Sprintf("%s%sCheckErr", strings.ToLower(fieldName[:1]), fieldName[1:])
	if !isPtr {
		if isOptional {
			format = `if reflect.DeepEqual(%s.%s,%s{}){
		err := %s.%s.Check()
		if err != nil {
			return err
		}
	}`
			return fmt.Sprintf(format, shortName, fieldName, structName, shortName, fieldName)

		} else {
			format = `if reflect.DeepEqual(%s.%s,%s{}){
		return fmt.Errorf("field %s is Required")
	}
	%s := %s.%s.Check()
	if %s != nil {
		return %s
	}`
			return fmt.Sprintf(format, shortName, fieldName, structName, fieldName, varName, shortName, fieldName, varName, varName)
		}
	} else {
		if isOptional {
			format = `if %s.%s!=nil{
		err := %s.%s.Check()
		if err != nil {
			return err
		}
	}`
			return fmt.Sprintf(format, shortName, fieldName, shortName, fieldName)

		} else {
			format = `if %s.%s==nil{
		return fmt.Errorf("field %s is Required")
	}
	%s := %s.%s.Check()
	if %s != nil {
		return %s
	}`
			return fmt.Sprintf(format, shortName, fieldName, fieldName, varName, shortName, fieldName, varName, varName)
		}
	}
}

func elemsBuilder(elemType string, cValue string) string {
	var elems string
	vs := strings.Split(cValue, ",")
	switch elemType {
	case Int8:
		elems = "[]int8{"
	case Int16:
		elems = "[]int16{"
	case Int32:
		elems = "[]int32{"
	case Int:
		elems = "[]int{"
	case Int64:
		elems = "[]int64{"
	case Uint8:
		elems = "[]uint8{"
	case Uint16:
		elems = "[]uint16{"
	case Uint32:
		elems = "[]uint32{"
	case Uint:
		elems = "[]uint{"
	case Uint64:
		elems = "[]uint64{"
	case Float32:
		elems = "[]float32{"
	case Float64:
		elems = "[]float64{"
	case String:
		elems = "[]string{"
	}
	for i := 0; i < len(vs); i++ {
		if elemType == String {
			elems += fmt.Sprintf("%q,", vs[i])
		} else {
			elems += fmt.Sprintf("%v,", vs[i])
		}
	}
	elems = strings.TrimRight(elems, ",") + "}"

	return elems
}
