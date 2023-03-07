package nut

import (
	"testing"
)

func Test_parse(t *testing.T) {
	opt := specimen{}
	_, err := parse(opt)
	t.Log("err:", err)
}
