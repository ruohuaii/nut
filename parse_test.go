package nut

import "testing"

func Test_parse(t *testing.T) {
	type data struct {
		Times int8 `nut:"between:-129,127"`
	}
	_, err := parse(data{Times: 2})
	t.Log("err:", err)
}
