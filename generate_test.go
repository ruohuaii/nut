package nut

import (
	"errors"
	"testing"
)

type MySelf struct {
	Name string `nut:"size:3"`
}

func Test_Generate(t *testing.T) {
	err := Generate(&MySelf{}, "generate_test.go", false)

	t.Log("generate_err:", err)
}

func Test_Generate_Check(t *testing.T) {
	self := &MySelf{
		Name: "Self-Denial",
	}

	t.Log("校验结果:", self.Check())
}

//Executing the "Test_Generate" method will generate this "Check" verification method
func (m *MySelf) Check() error {
	if len(m.Name) != 3 {
		return errors.New("the length of the Name field must be 3")
	}

	return nil
}
