package nut

import "fmt"

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
	format := `if %s.%s != %v {
     	return fmt.Errorf("the value of %s field should be less than %v")
	}`
	return fmt.Sprintf(format, shortName, fieldName, cValue, fieldName, cValue)
}

func ThrowCondLte(shortName, fieldName string, cValue any) string {
	format := `if %s.%s != %v {
     	return fmt.Errorf("the value of %s field should be less than or equal to %v")
	}`
	return fmt.Sprintf(format, shortName, fieldName, cValue, fieldName, cValue)
}

func ThrowCondGt(shortName, fieldName string, cValue any) string {
	format := `if %s.%s != %v {
     	return fmt.Errorf("the value of %s field should be greater than %v")
	}`
	return fmt.Sprintf(format, shortName, fieldName, cValue, fieldName, cValue)
}

func ThrowCondGte(shortName, fieldName string, cValue any) string {
	format := `if %s.%s != %v {
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
	case "string", "array":
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
