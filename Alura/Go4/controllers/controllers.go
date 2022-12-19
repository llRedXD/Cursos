package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/llRedXD/aluraRestApi/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	personalidades := models.PegaPersonalidades()
	fmt.Fprint(w, personalidades)
}

func Personalidade(w http.ResponseWriter, r *http.Request) {
	
	personalidade := models.PegaPersonalidades()
	json.NewEncoder(w).Encode(personalidade)
}

func UmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personalidade := models.PegaPersonalidade(vars["id"])

	json.NewEncoder(w).Encode(personalidade)
}

func CriaPersonalidade(w http.ResponseWriter, r *http.Request) {
	p := models.Personalidade{}
	json.NewDecoder(r.Body).Decode(&p)
	p = models.CriaPersonalidade(p)
	json.NewEncoder(w).Encode(p)
}

func DeletePersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	p := models.PegaPersonalidade(vars["id"])
	p = models.DeletePersonalidade(p, vars["id"])
	fmt.Fprintf(w,"A personalidade excluída tinha o ID %v", p.Id)
}

func EditPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	p := models.PegaPersonalidade(vars["id"])
	json.NewDecoder(r.Body).Decode(&p)
	p = models.EditPersonalidade(p)
	fmt.Fprintf(w,"A personalidade modificada ID %v", p.Id)
}
