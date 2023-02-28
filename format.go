package nut

import (
	"fmt"
	"strings"
)

func ThrowCondEq(shortName, fieldName string, cValue any) string {
	format := `if %s.%s == %v {
     	return fmt.Errorf("the value of %s field should not be equal to %v")
	}`
	return fmt.Sprintf(format, shortName, fieldName, cValue, fieldName, cValue)
}

func ThrowCondNeq(shortName, fieldName string, cValue any) string {
	format := `if %s.%s != %v {
     	return fmt.Errorf("the value of %s field should be equal to %v")
	}`
	return fmt.Sprintf(format, shortName, fieldName, cValue, fieldName, cValue)
}

func ThrowCondLt(shortName, fieldName string, cValue any) string {
	format := `if %s.%s < %v {
     	return fmt.Errorf("the value of %s field should be less than %v")
	}`
	return fmt.Sprintf(format, shortName, fieldName, cValue, fieldName, cValue)
}

func ThrowCondLte(shortName, fieldName string, cValue any) string {
	format := `if %s.%s <= %v {
     	return fmt.Errorf("the value of %s field should be less than or equal to %v")
	}`
	return fmt.Sprintf(format, shortName, fieldName, cValue, fieldName, cValue)
}

func ThrowCondGt(shortName, fieldName string, cValue any) string {
	format := `if %s.%s > %v {
     	return fmt.Errorf("the value of %s field should be greater than %v")
	}`
	return fmt.Sprintf(format, shortName, fieldName, cValue, fieldName, cValue)
}

func ThrowCondGte(shortName, fieldName string, cValue any) string {
	format := `if %s.%s >= %v {
     	return fmt.Errorf("the value of %s field should be greater than or equal to %v")
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
		elems += vs[i] + ","
	}
	elems = strings.TrimRight(elems, ",") + "}"
	varName := fmt.Sprintf("%sContains", strings.ToLower(fieldName))
	format := `var %s= %s
	for i:=0;i<len(%s);i++{
		if ArrayContains(%s.%s,%s[i]) {
			return fmt.Errorf("the value of the %s field must contain %s")
		}
	}`

	return fmt.Sprintf(format, varName, elems, varName, shortName, fieldName, varName, fieldName, cValue)
}

func ThrowCondExcluded(shortName, fieldName string, cValue string, elemType string) string {
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
		elems += vs[i] + ","
	}
	elems = strings.TrimRight(elems, ",") + "}"
	varName := fmt.Sprintf("%sExcluded", strings.ToLower(fieldName))
	format := `var %s = %s
	for i:=0;i<len(%s);i++{
		if !ArrayContains(%s.%s,%s[i]) {
			return fmt.Errorf("the value of the %s field cannot contain %s")
		}
	}`

	return fmt.Sprintf(format, varName, elems, varName, shortName, fieldName, varName, fieldName, cValue)
}

func ThrowCondIn(shortName, fieldName string, cValue string, elemType string) string {
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
	varName := fmt.Sprintf("%sIn", strings.ToLower(fieldName))
	format := `var %s = %s
	if !ArrayContains(%s,%s.%s) {
		return fmt.Errorf("the value of the %s field is not in %s")
	}`

	return fmt.Sprintf(format, varName, elems, varName, shortName, fieldName, fieldName, cValue)
}
