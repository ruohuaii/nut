package nut

type Field struct {
	Condition string
	//CondType  string
	//CondValue any
	//IsRequired bool
	//IsOptional bool
}

type Structure struct {
	FullName  string
	ShortName string
	Rules     map[string][]Field
	//Eq              []*Field
	//Neq             []*Field
	//Gt              []*Field
	//Gte             []*Field
	//Lt              []*Field
	//Lte             []*Field
	//Between         []*Field
	//Size            []*Field
	//In              []*Field
	//Contains        []*Field
}
