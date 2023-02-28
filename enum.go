package nut

//Nut tag name
const Nut = "nut"

//Zero numeral base value
const Zero = 0

//Some conditions name
const (
	Eq       = "eq"
	Neq      = "neq"
	Gt       = "gt"
	Gte      = "gte"
	Lt       = "lt"
	Lte      = "lte"
	Between  = "between"
	In       = "in"
	Contains = "contains"
	Excluded = "excluded"
	Size     = "size"
	Regexp   = "regexp"

	//Type array or slice element type
	Type = "type"
)

//Some conditions set
var (
	NumeralCondSet = [10]string{Eq, Neq, Lt, Lte, Gt, Gte, Between, In, Type, Size}
	StringCondSet  = [11]string{Eq, Neq, Lt, Lte, Gt, Gte, Between, In, Type, Regexp, Size}
	ArrayCondSet   = [4]string{Contains, Excluded, Type, Size}
	BoolCondSet    = [1]string{Eq}
)

const (
	Int8    = "int8"
	Int16   = "int16"
	Int32   = "int32"
	Int     = "int"
	Int64   = "int64"
	Uint8   = "uint8"
	Uint16  = "uint16"
	Uint32  = "uint32"
	Uint64  = "uint64"
	Float32 = "float32"
	Float64 = "float64"
	String  = "string"
)
