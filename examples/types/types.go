package types

import (
	"errors"
	"strconv"

	. "github.com/ruohuaii/nut"
)

type (
	Person struct {
		Hobby             []string `nut:"optional;size:3;contains:sing,rap,basketball;type:string"`
		Nature            string   `nut:"optional;size:3,30"`
		FavoriteNumber    string   `nut:"optional;type:uint8"`
		FavoriteSubject   string   `nut:"optional;in:Math,English,Physics"`
		DisgustingSubject []string `nut:"optional;excluded:Math,Physics;type:string"`
		Jenny             *Jenny   `nut:"optional;summary"`
		Robert            *Robert  `nut:"required;summary"`
	}

	Jenny struct {
		Name string `nut:"required;size:5"`
		Age  uint8  `nut:"required;gt:10"`
	}

	Robert struct {
		Name string `nut:"required;size:6"`
		Age  uint8  `nut:"required;gte:10"`
	}
)

func (p *Person) Summary() map[string]map[string]string {
	return map[string]map[string]string{
		"Hobby":  {"size": "it is hobby"},
		"Nature": {"size": "it is nature"},
	}
}

func (j *Jenny) Summary() map[string]map[string]string {
	return map[string]map[string]string{
		"Hobby": {"size": "你在干什么啊"},
	}
}

func (r *Robert) Summary() map[string]map[string]string {
	return map[string]map[string]string{
		"Hobby": {"size": "你在干什么啊"},
	}
}
func (j *Jenny) Check() error {
	if j.Age <= 10 {
		return errors.New("the value of Age field should be greater than 10")
	}
	if len(j.Name) != 5 {
		return errors.New("the length of the Name field must be 5")
	}

	return nil
}
func (r *Robert) Check() error {
	if r.Age < 10 {
		return errors.New("the value of Age field should be greater than or equal to 10")
	}
	if len(r.Name) != 6 {
		return errors.New("the length of the Name field must be 6")
	}

	return nil
}
func (p *Person) Check() error {
	var disgustingSubjectContains = []string{"Math", "Physics"}
	for i := 0; i < len(disgustingSubjectContains); i++ {
		if ArrayContains(p.DisgustingSubject, disgustingSubjectContains[i]) {
			return errors.New("the value of the DisgustingSubject field cannot contain Math,Physics")
		}
	}
	_, favoriteNumberParseErr := strconv.ParseUint(p.FavoriteNumber, 10, 64)
	if favoriteNumberParseErr != nil {
		return errors.New("the value of the FavoriteNumber field is wrong")
	}
	if p.FavoriteNumber != "" {
		_, favoriteNumberParseErr := strconv.ParseUint(p.FavoriteNumber, 10, 64)
		if favoriteNumberParseErr != nil {
			return errors.New("the value of the FavoriteNumber field is wrong")
		}
	}
	var favoriteSubjectIn = []string{"Math", "English", "Physics"}
	if !ArrayContains(favoriteSubjectIn, p.FavoriteSubject) {
		return errors.New("the value of the FavoriteSubject field should be one of Math,English,Physics")
	}
	if p.FavoriteSubject != "" {
		var favoriteSubjectIn = []string{"Math", "English", "Physics"}
		if !ArrayContains(favoriteSubjectIn, p.FavoriteSubject) {
			return errors.New("the value of the FavoriteSubject field should be one of Math,English,Physics")
		}
	}
	if len(p.Hobby) != 3 {
		return errors.New("it is hobby")
	}
	var hobbyExcluded = []string{"sing", "rap", "basketball"}
	for i := 0; i < len(hobbyExcluded); i++ {
		if !ArrayContains(p.Hobby, hobbyExcluded[i]) {
			return errors.New("the value of the Hobby field must contain sing,rap,basketball")
		}
	}
	if p.Nature != "" {
		if len(p.Nature) < 3 || len(p.Nature) > 30 {
			return errors.New("it is nature")
		}
	}
	if len(p.Nature) < 3 || len(p.Nature) > 30 {
		return errors.New("it is nature")
	}
	if p.Jenny != nil {
		err := p.Jenny.Check()
		if err != nil {
			return err
		}
	}
	if p.Robert == nil {
		return errors.New("field Robert is Required")
	}
	robertCheckErr := p.Robert.Check()
	if robertCheckErr != nil {
		return robertCheckErr
	}

	return nil
}
