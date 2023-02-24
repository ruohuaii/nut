package nut

type Field struct {
	Name       string
	Value      any
	IsRequired bool
	IsOptional bool
}

type Structure struct {
	StructFullName  string
	StructShortName string
	Eq              []*Field
	Neq             []*Field
	Gt              []*Field
	Gte             []*Field
	Lt              []*Field
	Lte             []*Field
	Between         []*Field
	In              []*Field
	Contains        []*Field
}
