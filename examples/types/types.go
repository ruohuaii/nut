package types

import (
	"fmt"
	"math"
	"reflect"
	"strconv"

	. "github.com/ruohuaii/nut"
)

type (
	Person struct {
		Hobby             []string `nut:"size:3;contains:sing,rap,basketball;type:string"`
		Nature            string   `nut:"size:3,30"`
		FavoriteNumber    string   `nut:"type:uint8"`
		FavoriteSubject   string   `nut:"in:Math,English,Physics"`
		DisgustingSubject []string `nut:"excluded:Math,Physics;type:string"`
		Jenny             *Jenny   `nut:"optional"`
		Robert            Robert   `nut:"required"`
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
	var disgustingsubjectContains = []string{"Math", "Physics"}
	for i := 0; i < len(disgustingsubjectContains); i++ {
		if ArrayContains(p.DisgustingSubject, disgustingsubjectContains[i]) {
			return fmt.Errorf("the value of the DisgustingSubject field cannot contain Math,Physics")
		}
	}
	favoriteNumberVal, err := strconv.ParseUint(p.FavoriteNumber, 10, 64)
	if err != nil {
		return fmt.Errorf("the value of the FavoriteNumber field is wrong")
	}
	if favoriteNumberVal > math.MaxUint8 || favoriteNumberVal < 0 {
		return fmt.Errorf("the value of the FavoriteNumber field overflowed")
	}
	var favoriteSubjectIn = []string{"Math", "English", "Physics"}
	if !ArrayContains(favoriteSubjectIn, p.FavoriteSubject) {
		return fmt.Errorf("the value of the FavoriteSubject field should be one of Math,English,Physics")
	}
	if len(p.Hobby) != 3 {
		return fmt.Errorf("the length of the Hobby field must be 3")
	}
	var hobbyExcluded = []string{"sing", "rap", "basketball"}
	for i := 0; i < len(hobbyExcluded); i++ {
		if !ArrayContains(p.Hobby, hobbyExcluded[i]) {
			return fmt.Errorf("the value of the Hobby field must contain sing,rap,basketball")
		}
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
