package nut

import (
	"fmt"
	"regexp"
)

type Woman struct {
	Age     int     `nut:"eq:18"`
	Weight  int     `nut:"neq:120"`
	Height  int     `nut:"between:170,180"`
	Belt    int8    `nut:"size:1"`
	Like    []int8  `nut:"size:1,5;excluded:2,3,4;type:int8"`
	Study   string  `nut:"regexp:/d"`
	Boot    string  `nut:"in:l1,l2,l3;type:string"`
	Husband Husband `nut:"skip"`
}

type Husband struct {
	Name string `nut:"size:10,20"`
	Age  int    `nut:"eq:20"`
}

func (w *Woman) Check() error {
	if w.Age != 18 {
		return fmt.Errorf("the value of Age field should be equal to 18")
	}
	if len(fmt.Sprint(w.Belt)) != 1 {
		return fmt.Errorf("the length of the Belt field must be 1")
	}
	var bootIn = []string{"l1", "l2", "l3"}
	if !ArrayContains(bootIn, w.Boot) {
		return fmt.Errorf("the value of the Boot field is not in l1,l2,l3")
	}
	if w.Height < 170 || w.Height > 180 {
		return fmt.Errorf("the value of Height field should be between 170 and 180")
	}
	if len(w.Like) < 1 || len(w.Like) > 5 {
		return fmt.Errorf("the length of Like field value should be between 1 and 5")
	}
	var likeContains = []int8{2, 3, 4}
	for i := 0; i < len(likeContains); i++ {
		if ArrayContains(w.Like, likeContains[i]) {
			return fmt.Errorf("the value of the Like field must contain 2,3,4")
		}
	}
	studyRegexp, err := regexp.Compile("/d")
	if err != nil {
		return err
	}
	if !studyRegexp.MatchString(w.Study) {
		return fmt.Errorf("the value of the Study field does not conform to the regular rules")
	}
	if w.Weight == 120 {
		return fmt.Errorf("the value of Weight field should not be equal to 120")
	}

	return nil
}
