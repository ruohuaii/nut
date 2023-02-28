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

func ThrowCondSize(shortName, fieldName string, left, right any, fieldType string) string {
	var format string
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

	return fmt.Sprintf(format, shortName, fieldName, left, shortName, fieldName, right,
		fieldName, left, right)
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
	format := `var shouldContains = %s
	for i:=0;i<len(shouldContains);i++{
		if ArrayContains(%s.%s,shouldContains[i]) {
			return fmt.Errorf("the value of the %s field must contain %s")
		}
	}`

	return fmt.Sprintf(format, elems, shortName, fieldName, fieldName, cValue)
}

func ThrowCondNContain(shortName, fieldName string, cValue string, elemType string) string {
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
	format := `var shouldContains = %s
	for i:=0;i<len(shouldContains);i++{
		if !ArrayContains(%s.%s,shouldContains[i]) {
			return fmt.Errorf("the value of the %s field cannot contain %s")
		}
	}`

	return fmt.Sprintf(format, elems, shortName, fieldName, fieldName, cValue)
}
