package nut

//Nut is the structure tag name
const Nut = "nut"

//Zero is the underlying value of the numeric type
const Zero = 0

//Semicolon is the tag separator
const Semicolon = ";"

const StringNull = ""

//Some condition names
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

	//Type is used to describe the type of the field or the type of the array element
	Type = "type"

	Optional = "optional"
	Required = "required"

	//Associate  Field Conditions
	Associate = "associate"
)

//Some set of conditions
var (
	NumeralCondSet = [11]string{Eq, Neq, Lt, Lte, Gt, Gte, Between, In, Type, Size, Associate}
	StringCondSet  = [7]string{Eq, Neq, In, Type, Regexp, Size, Associate}
	ArrayCondSet   = [5]string{Contains, Excluded, Type, Size, Associate}
	BoolCondSet    = [2]string{Eq, Associate}
	StructCondSet  = [2]string{Required, Optional}
	Types          = [13]string{Int8, Int16, Int32, Int, Int64, Uint8, Uint16, Uint32, Uint, Uint64, Float32, Float64, String}
)

//Some types name

const (
	Int8    = "int8"
	Int16   = "int16"
	Int32   = "int32"
	Int     = "int"
	Int64   = "int64"
	Uint8   = "uint8"
	Uint16  = "uint16"
	Uint32  = "uint32"
	Uint    = "uint"
	Uint64  = "uint64"
	Float32 = "float32"
	Float64 = "float64"
	String  = "string"
)

const StructField = "StructField"
