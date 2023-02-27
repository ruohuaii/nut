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
{{range $k,$v:=.Rules}}
{{range $v}}
	if {{.Condition}} {
		return errors.New("出错啦!")
	}
{{end}}
{{end}}
}
`
