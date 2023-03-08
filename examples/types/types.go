package types

import (
	"fmt"
	"reflect"
)

type (
	Person struct {
		Hobby          string `nut:"size:1,10"`
		Nature         string `nut:"size:3,30"`
		FavoriteNumber uint8  `nut:"between:0,10"`
		Jenny          *Jenny `nut:"optional"`
		Robert         Robert `nut:"required"`
	}

	Jenny struct {
		Name string `nut:"size:5"`
		Age  uint8  `nut:"gt:10"`
	}

	Robert struct {
		Name string `nut:"size:6"`
		Age  uint8  `nut:"gte:10"`
	}
)

func (j *Jenny) Check() error {
	if j.Age <= 10 {
		return fmt.Errorf("the value of Age field should be greater than 10")
	}
	if len(j.Name) != 5 {
		return fmt.Errorf("the length of the Name field must be 5")
	}

	return nil
}
func (r *Robert) Check() error {
	if r.Age < 10 {
		return fmt.Errorf("the value of Age field should be greater than or equal to 10")
	}
	if len(r.Name) != 6 {
		return fmt.Errorf("the length of the Name field must be 6")
	}

	return nil
}
func (p *Person) Check() error {
	if p.FavoriteNumber < 0 || p.FavoriteNumber > 10 {
		return fmt.Errorf("the value of FavoriteNumber field should be between 0 and 10")
	}
	if len(p.Hobby) < 1 || len(p.Hobby) > 10 {
		return fmt.Errorf("the length of Hobby field value should be between 1 and 10")
	}
	if len(p.Nature) < 3 || len(p.Nature) > 30 {
		return fmt.Errorf("the length of Nature field value should be between 3 and 30")
	}
	if p.Jenny != nil {
		err := p.Jenny.Check()
		if err != nil {
			return err
		}
	}
	if reflect.DeepEqual(p.Robert, Robert{}) {
		return fmt.Errorf("field Robert is Required")
	}
	robertCheckErr := p.Robert.Check()
	if robertCheckErr != nil {
		return robertCheckErr
	}

	return nil
}
