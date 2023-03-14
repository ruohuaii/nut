package nut

import (
	"fmt"
	"strings"
)

func ThrowCondEq(shortName, fieldName string, cValue any, isString bool, summary string) string {
	var format string
	if summary == "" {
		if !isString {
			format = `if %s.%s == %v {
     	return errors.New("the value of %s field should not be equal to %v")
	}`
		} else {
			format = `if %s.%s == "%v" {
     	return errors.New("the value of %s field should not be equal to %v")
	}`
		}

		return fmt.Sprintf(format, shortName, fieldName, cValue, fieldName, cValue)
	} else {
		if !isString {
			format = `if %s.%s == %v {
     	return errors.New(%q)
	}`
		} else {
			format = `if %s.%s == "%v" {
     	return errors.New(%q)
	}`
		}

		return fmt.Sprintf(format, shortName, fieldName, cValue, summary)
	}

}

func ThrowCondNeq(shortName, fieldName string, cValue any, isString bool, summary string) string {
	var format string
	if summary == "" {
		if !isString {
			format = `if %s.%s != %v {
     	return errors.New("the value of %s field should be equal to %v")
	}`
		} else {
			format = `if %s.%s != "%v" {
     	return errors.New("the value of %s field should be equal to %v")
	}`
		}

		return fmt.Sprintf(format, shortName, fieldName, cValue, fieldName, cValue)
	} else {
		if !isString {
			format = `if %s.%s != %v {
     	return errors.New(%q)
	}`
		} else {
			format = `if %s.%s != "%v" {
     	return errors.New(%q)
	}`
		}

		return fmt.Sprintf(format, shortName, fieldName, cValue, summary)
	}

}

func ThrowCondLt(shortName, fieldName string, cValue any, summary string) string {
	if summary == "" {
		format := `if %s.%s < %v {
		return errors.New("the value of %s field should be greater than or equal to %v")
	}`
		return fmt.Sprintf(format, shortName, fieldName, cValue, fieldName, cValue)
	} else {
		format := `if %s.%s < %v {
		return errors.New(%q)
	}`
		return fmt.Sprintf(format, shortName, fieldName, cValue, summary)
	}

}

func ThrowCondLte(shortName, fieldName string, cValue any, summary string) string {
	if summary == "" {
		format := `if %s.%s <= %v {
		return errors.New("the value of %s field should be greater than %v")
	}`
		return fmt.Sprintf(format, shortName, fieldName, cValue, fieldName, cValue)
	} else {
		format := `if %s.%s <= %v {
		return errors.New(%q)
	}`
		return fmt.Sprintf(format, shortName, fieldName, cValue, summary)
	}
}

func ThrowCondGt(shortName, fieldName string, cValue any, summary string) string {
	if summary == "" {
		format := `if %s.%s > %v {
		return errors.New("the value of %s field should be less than or equal to %v")
	}`
		return fmt.Sprintf(format, shortName, fieldName, cValue, fieldName, cValue)
	} else {
		format := `if %s.%s > %v {
		return errors.New(%q)
	}`
		return fmt.Sprintf(format, shortName, fieldName, cValue, summary)
	}
}

func ThrowCondGte(shortName, fieldName string, cValue any, summary string) string {
	if summary == "" {
		format := `if %s.%s >= %v {
		return errors.New("the value of %s field should be less than %v")
	}`
		return fmt.Sprintf(format, shortName, fieldName, cValue, fieldName, cValue)
	} else {
		format := `if %s.%s >= %v {
		return errors.New(%q)
	}`
		return fmt.Sprintf(format, shortName, fieldName, cValue, summary)
	}
}

func ThrowCondBetween(shortName, fieldName string, left, right any, summary string) string {
	if summary == "" {
		format := `if %s.%s < %v || %s.%s > %v {
     	return errors.New("the value of %s field should be between %v and %v")
	}`
		return fmt.Sprintf(format, shortName, fieldName, left, shortName, fieldName, right,
			fieldName, left, right)
	} else {
		format := `if %s.%s < %v || %s.%s > %v {
     	return errors.New(%q)
	}`
		return fmt.Sprintf(format, shortName, fieldName, left, shortName, fieldName, right, summary)
	}

}

func ThrowCondRegexp(shortName, fieldName string, matchRule string, summary string) string {
	varName := fmt.Sprintf("%sRegexp", strings.ToLower(fieldName))
	if summary == "" {
		format := `%s,err := regexp.Compile(%q)
	if err != nil {
		return err
	}
	if !%s.MatchString(%s.%s) {
		return errors.New("the value of the %s field does not conform to the regular rules")
	}`
		return fmt.Sprintf(format, varName, matchRule, varName, shortName, fieldName, fieldName)
	} else {
		format := `%s,err := regexp.Compile(%q)
	if err != nil {
		return err
	}
	if !%s.MatchString(%s.%s) {
		return errors.New(%q)
	}`
		return fmt.Sprintf(format, varName, matchRule, varName, shortName, fieldName, summary)
	}
}

func ThrowCondSize(shortName, fieldName string, cvs []string, fieldType string, summary string) string {
	var format string
	if summary == "" {
		if len(cvs) == 2 {
			switch fieldType {
			case "string", "array", "slice":
				format = `if len(%s.%s) < %v || len(%s.%s) > %v {
		return errors.New("the length of %s field value should be between %v and %v")
	}`
			case "int8", "int16", "int32", "int", "int64",
				"uint8", "uint16", "uint32", "uint", "uint64",
				"float32", "float64":
				format = `if len(fmt.Sprint(%s.%s)) < %v || len(fmt.Sprint(%s.%s)) > %v {
		return errors.New("the length of %s field value should be between %v and %v")
	}`
			}
			left, right := cvs[0], cvs[1]
			return fmt.Sprintf(format, shortName, fieldName, left, shortName, fieldName, right,
				fieldName, left, right)
		} else {
			switch fieldType {
			case "string", "array", "slice":
				format = `if len(%s.%s) != %s {
		return errors.New("the length of the %s field must be %s")
	}`
			case "int8", "int16", "int32", "int", "int64",
				"uint8", "uint16", "uint32", "uint", "uint64",
				"float32", "float64":
				format = `if len(fmt.Sprint(%s.%s)) != %s {
		return errors.New("the length of the %s field must be %s")
	}`
			}
			length := cvs[0]
			return fmt.Sprintf(format, shortName, fieldName, length, fieldName, length)
		}
	} else {
		if len(cvs) == 2 {
			switch fieldType {
			case "string", "array", "slice":
				format = `if len(%s.%s) < %v || len(%s.%s) > %v {
		return errors.New(%q)
	}`
			case "int8", "int16", "int32", "int", "int64",
				"uint8", "uint16", "uint32", "uint", "uint64",
				"float32", "float64":
				format = `if len(fmt.Sprint(%s.%s)) < %v || len(fmt.Sprint(%s.%s)) > %v {
		return errors.New(%q)
	}`
			}
			left, right := cvs[0], cvs[1]
			return fmt.Sprintf(format, shortName, fieldName, left, shortName, fieldName, right, summary)
		} else {
			switch fieldType {
			case "string", "array", "slice":
				format = `if len(%s.%s) != %s {
		return errors.New(%q)
	}`
			case "int8", "int16", "int32", "int", "int64",
				"uint8", "uint16", "uint32", "uint", "uint64",
				"float32", "float64":
				format = `if len(fmt.Sprint(%s.%s)) != %s {
		return errors.New(%q)
	}`
			}
			length := cvs[0]
			return fmt.Sprintf(format, shortName, fieldName, length, summary)
		}
	}
}

func ThrowCondContains(shortName, fieldName string, cValue string, elemType string, summary string) string {
	elems := elemsBuilder(elemType, cValue)
	varName := fmt.Sprintf("%s%sContains", strings.ToLower(fieldName[:1]), fieldName[1:])
	if summary == "" {
		format := `var %s= %s
	for i:=0;i<len(%s);i++{
		if ArrayContains(%s.%s,%s[i]) {
			return errors.New("the value of the %s field cannot contain %s")
		}
	}`

		return fmt.Sprintf(format, varName, elems, varName, shortName, fieldName, varName, fieldName, cValue)
	} else {
		format := `var %s= %s
	for i:=0;i<len(%s);i++{
		if ArrayContains(%s.%s,%s[i]) {
			return errors.New(%q)
		}
	}`

		return fmt.Sprintf(format, varName, elems, varName, shortName, fieldName, varName, summary)
	}

}

func ThrowCondExcluded(shortName, fieldName string, cValue string, elemType string, summary string) string {
	elems := elemsBuilder(elemType, cValue)
	varName := fmt.Sprintf("%s%sExcluded", strings.ToLower(fieldName[:1]), fieldName[1:])
	if summary == "" {
		format := `var %s = %s
	for i:=0;i<len(%s);i++{
		if !ArrayContains(%s.%s,%s[i]) {
			return errors.New("the value of the %s field must contain %s")
		}
	}`

		return fmt.Sprintf(format, varName, elems, varName, shortName, fieldName, varName, fieldName, cValue)
	} else {
		format := `var %s = %s
	for i:=0;i<len(%s);i++{
		if !ArrayContains(%s.%s,%s[i]) {
			return errors.New(%q)
		}
	}`

		return fmt.Sprintf(format, varName, elems, varName, shortName, fieldName, varName, summary)
	}
}

func ThrowCondIn(shortName, fieldName string, cValue string, elemType string, summary string) string {
	elems := elemsBuilder(elemType, cValue)
	varName := fmt.Sprintf("%s%sIn", strings.ToLower(fieldName[:1]), fieldName[1:])
	if summary == "" {
		format := `var %s = %s
	if !ArrayContains(%s,%s.%s) {
		return errors.New("the value of the %s field should be one of %s")
	}`

		return fmt.Sprintf(format, varName, elems, varName, shortName, fieldName, fieldName, cValue)
	} else {
		format := `var %s = %s
	if !ArrayContains(%s,%s.%s) {
		return errors.New(%q)
	}`

		return fmt.Sprintf(format, varName, elems, varName, shortName, fieldName, summary)
	}
}

func ThrowCondType(shortName, fieldName string, elemType string, summary string) string {
	var format string
	varName := fmt.Sprintf("%s%sParseErr", strings.ToLower(fieldName[:1]), fieldName[1:])
	if summary == "" {
		switch elemType {
		case Int8:
			format = `_,%s := strconv.ParseInt(%s.%s,10,64)
	if %s != nil {
		return errors.New("the value of the %s field is wrong")
	}`
		case Int16:
			format = `_,%s := strconv.ParseInt(%s.%s,10,64)
	if %s != nil {
		return errors.New("the value of the %s field is wrong")
	}`
		case Int32:
			format = `_,%s := strconv.ParseInt(%s.%s,10,64)
	if %s != nil {
		return errors.New("the value of the %s field is wrong")
	}`
		case Int:
			format = `_,%s := strconv.ParseInt(%s.%s,10,64)
	if %s != nil {
		return errors.New("the value of the %s field is wrong")
	}`
		case Int64:
			format = `_,%s := strconv.ParseInt(%s.%s,10,64)
	if %s != nil {
		return errors.New("the value of the %s field is wrong")
	}`
		case Uint8:
			format = `_,%s := strconv.ParseUint(%s.%s,10,64)
	if %s != nil {
		return errors.New("the value of the %s field is wrong")
	}`
		case Uint16:
			format = `_,%s := strconv.ParseUint(%s.%s,10,64)
	if %s != nil {
		return errors.New("the value of the %s field is wrong")
	}`
		case Uint32:
			format = `_,%s := strconv.ParseUint(%s.%s,10,64)
	if %s != nil {
		return errors.New("the value of the %s field is wrong")
	}`
		case Uint64:
			format = `_,%s := strconv.ParseUint(%s.%s,10,64)
	if %s != nil {
		return errors.New("the value of the %s field is wrong")
	}`
		case Float32:
			format = `_,%s := strconv.ParseFloat(%s.%s,32)
	if %s != nil {
		return errors.New("the value of the %s field is wrong")
	}`
		case Float64:
			format = `_,%s := strconv.ParseFloat(%s.%s,64)
	if %s != nil {
		return errors.New("the value of the %s field is wrong")
	}`
		}

		return fmt.Sprintf(format, varName, shortName, fieldName, varName, fieldName)
	} else {
		switch elemType {
		case Int8:
			format = `_,%s := strconv.ParseInt(%s.%s,10,64)
	if %s != nil {
		return errors.New(%q)
	}`
		case Int16:
			format = `_,%s := strconv.ParseInt(%s.%s,10,64)
	if %s != nil {
		return errors.New(%q)
	}`
		case Int32:
			format = `_,%s := strconv.ParseInt(%s.%s,10,64)
	if %s != nil {
		return errors.New(%q)
	}`
		case Int:
			format = `_,%s := strconv.ParseInt(%s.%s,10,64)
	if %s != nil {
		return errors.New(%q)
	}`
		case Int64:
			format = `_,%s := strconv.ParseInt(%s.%s,10,64)
	if %s != nil {
		return errors.New(%q)
	}`
		case Uint8:
			format = `_,%s := strconv.ParseUint(%s.%s,10,64)
	if %s != nil {
		return errors.New(%q)
	}`
		case Uint16:
			format = `_,%s := strconv.ParseUint(%s.%s,10,64)
	if %s != nil {
		return errors.New(%q)
	}`
		case Uint32:
			format = `_,%s := strconv.ParseUint(%s.%s,10,64)
	if %s != nil {
		return errors.New(%q)
	}`
		case Uint64:
			format = `_,%s := strconv.ParseUint(%s.%s,10,64)
	if %s != nil {
		return errors.New(%q)
	}`
		case Float32:
			format = `_,%s := strconv.ParseFloat(%s.%s,32)
	if %s != nil {
		return errors.New(%q)
	}`
		case Float64:
			format = `_,%s := strconv.ParseFloat(%s.%s,64)
	if %s != nil {
		return errors.New(%q)
	}`
		}

		return fmt.Sprintf(format, varName, shortName, fieldName, varName, summary)
	}
}

func ThrowCondStruct(shortName, fieldName, structName string, isOptional bool, isPtr bool, summary string) string {
	var format string
	varName := fmt.Sprintf("%s%sCheckErr", strings.ToLower(fieldName[:1]), fieldName[1:])
	if summary == "" {
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
		return errors.New("field %s is Required")
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
		return errors.New("field %s is Required")
	}
	%s := %s.%s.Check()
	if %s != nil {
		return %s
	}`
				return fmt.Sprintf(format, shortName, fieldName, fieldName, varName, shortName, fieldName, varName, varName)
			}
		}
	} else {
		if !isPtr {
			if isOptional {
				format = `if reflect.DeepEqual(%s.%s,%s{}){
		err := %s.%s.Check()
		if err != nil {
			return errors.New(%q)
		}
	}`
				return fmt.Sprintf(format, shortName, fieldName, structName, shortName, fieldName, summary)

			} else {
				format = `if reflect.DeepEqual(%s.%s,%s{}){
		return errors.New(%q)
	}
	%s := %s.%s.Check()
	if %s != nil {
		return errors.New(%q)
	}`
				return fmt.Sprintf(format, shortName, fieldName, structName, summary, varName, shortName, fieldName, varName, summary)
			}
		} else {
			if isOptional {
				format = `if %s.%s!=nil{
		err := %s.%s.Check()
		if err != nil {
			return errors.New(%q)
		}
	}`
				return fmt.Sprintf(format, shortName, fieldName, shortName, summary)

			} else {
				format = `if %s.%s==nil{
		return errors.New(%q)
	}
	%s := %s.%s.Check()
	if %s != nil {
		return errors.New(%q)
	}`
				return fmt.Sprintf(format, shortName, fieldName, summary, varName, shortName, fieldName, varName, summary)
			}
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
