package nut

import (
	"testing"
)

func Test_parse(t *testing.T) {
	opt := Structure{}
	_, err := Parse(opt)
	t.Log("err:", err)
}
