package nut

import "reflect"

type specimen struct {
	FullName   string
	ShortName  string
	Conditions map[string]Rules
}

type Relation struct {
	Self      string
	Associate string
}

type Rules struct {
	SelfRules      []Condition
	AssociateRules []Condition
}

type Condition struct {
	Rule  string
	FType string
}

type FieldStruct struct {
	FieldName  string
	Type       reflect.Type
	Value      reflect.Value
	IsOptional bool
	IsPtr      bool
	HasSummary bool
}
