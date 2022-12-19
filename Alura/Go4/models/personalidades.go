package models

import (
	"github.com/llRedXD/aluraRestApi/db"
)

type Personalidade struct {
	Id       int    `json:"id"`
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

func PegaPersonalidades() []Personalidade {
	db := db.ConectDb()
	p := []Personalidade{}
	db.Find(&p)
	return p
}

func PegaPersonalidade(id string) Personalidade {
	db := db.ConectDb()
	p := Personalidade{}
	db.Find(&p, id)
	return p
}

func CriaPersonalidade(p Personalidade) Personalidade {
	db := db.ConectDb()
	db.Create(&p)
	return p
}

func DeletePersonalidade(p Personalidade, id string) Personalidade {
	db := db.ConectDb()
	db.Delete(&p, id)
	return p
}

func EditPersonalidade(p Personalidade) Personalidade {
	db := db.ConectDb()
	db.Save(&p)
	return p
}

