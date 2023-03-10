package nut

import (
	"fmt"
	"text/template"
)

var tpl *template.Template

func init() {
	var err error
	tpl, err = template.New("NutVT").Parse(StdNutVT)
	if err != nil {
		panic(fmt.Sprintf("failed to parse template:%v", err))
	}
}

var StdNutVT = `func ({{.ShortName}} {{.FullName}}) Check() error {
{{range .Conditions}}{{range .SelfRules}}	{{.Rule}}
{{end}}{{end}}
	return nil
}
`
