package nut

import (
	"testing"
)

func Test_Generate(t *testing.T) {
	_, err := Generate(Woman{
		Age: 12,
	}, "woman_example.go")

	t.Log("generate_err:", err)
}

func Test_Generate_Check(t *testing.T) {
	woman := &Woman{
		Age:    18,
		Weight: 110,
		Height: 170,
		Belt:   1,
	}

	t.Log("校验结果:", woman.Check())
}
