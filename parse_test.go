package nut

import "testing"

func Test_parse(t *testing.T) {
	type Name struct {
		Name string `nut:"size:21"`
	}
	type data struct {
		Times *int8 `nut:"between:-128,127"`
		Name  *Name
		Boy   bool `nut:"eq:true"`
	}
	var a int8 = 2
	_, err := parse(data{Times: &a, Name: &Name{Name: "Hh"}, Boy: true})
	t.Log("err:", err)
}
