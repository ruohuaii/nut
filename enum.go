package nut

const Nut = "nut"

const (
	Require  = "require"
	Optional = "optional"
)

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
	Regexp   = "regexp"
)

var conditions = [9]string{Eq, Neq, Gt, Gte, Lt, Lte, Between, In, Contains}
