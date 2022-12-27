package models

import (
	"gorm.io/gorm"
	"gopkg.in/validator.v2"
)

type Aluno struct {
	gorm.Model
	Nome string `json:"nome" validate:"nonzero"`
	Cpf  string `json:"cpf"  validate:"len=11, regexp=^[0-9]*$"`
	Rg   string `json:"rg"   validate:"len=9,  regexp=^[0-9]*$"`
}

var Data []Aluno


func (a *Aluno) Validate() error {
	return validator.Validate(a)
}