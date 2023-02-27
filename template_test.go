package nut

import (
	"bytes"
	"os"
	"testing"
	"text/template"
)

type t_tpl struct {
	Type string
	Name string
}

func Test_Template(t *testing.T) {
	tpl, err := template.New("test").Parse(StdNutVT)
	if err != nil {
		panic("模板解析:" + err.Error())
	}

	opt := Structure{
		FullName:  "*Suit",
		ShortName: "s",
		Rules: map[string][]Field{
			"Name": {
				{
					Condition: "len(s.Name)>5",
				},
			},
		},
	}
	writer := bytes.Buffer{}
	err = tpl.Execute(&writer, opt)
	if err != nil {
		panic("模板填充:" + err.Error())
	}

	fd, err := os.OpenFile("main.go", os.O_APPEND, os.ModePerm)
	if err != nil {
		panic("打开文件:" + err.Error())
	}

	_, err = fd.Write(writer.Bytes())
	if err != nil {
		panic("写入文件:" + err.Error())
	}
}
