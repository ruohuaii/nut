package nut

import "fmt"

type Woman struct {
	Age    int `nut:"eq:18"`
	Weight int `nut:"neq:120"`
	Height int `nut:"between:170,180"`
	Belt   int `nut:"size:1,3"`
}

func (w *Woman) Check() error {
	if w.Age != 18 {
		return fmt.Errorf("the value of Age field should be equal to 18")
	}
	if len(fmt.Sprint(w.Belt)) < 1 || len(fmt.Sprint(w.Belt)) > 3 {
		return fmt.Errorf("the length of Belt field value should be between 90 and 110")
	}
	if w.Height < 170 || w.Height > 180 {
		return fmt.Errorf("the value of Height field should be between 170 and 180")
	}
	if w.Weight == 120 {
		return fmt.Errorf("the value of Weight field should not be equal to 120")
	}

	return nil
}
