package nut

type specimen struct {
	FullName   string
	ShortName  string
	Conditions map[string][]Condition
	Associates map[string][]Relation
}

type Relation struct {
	Self      string
	Associate []string
}

type Condition struct {
	Description string
}
