package nut

const Nut = "nut"

const Zero = 0

const (
	Require  = "require"
	Optional = "optional"
)

const (
	Null     = ""
	Eq       = "eq"
	Neq      = "neq"
	Gt       = "gt"
	Gte      = "gte"
	Lt       = "lt"
	Lte      = "lte"
	Between  = "between"
	In       = "in"
	Contains = "contains"
	Size     = "size"
	Regexp   = "regexp"
)

var (
	SingleTypeCondSet = [11]string{Eq, Neq, Gt, Gte, Lt, Lte, Between, In, Contains, Size}
	ArrayCondSet      = [4]string{Size, In, Contains}
	BoolCondSet       = [1]string{Eq}
)
