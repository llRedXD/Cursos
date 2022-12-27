package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/llRedXD/Go6/controllers"
	"github.com/llRedXD/Go6/db"
	"github.com/llRedXD/Go6/models"
	"github.com/stretchr/testify/assert"
)

func SetupDasRotasDeTeste() *gin.Engine {
	// Setup das rotas de teste
	r := gin.Default()
	return r
}

func SetupSaudacao() *httptest.ResponseRecorder {
	// Setup da Requisição Saudação
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacoes)
	req, _ := http.NewRequest("GET", "/ga", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestSaudacaoStatusCode(t *testing.T) {
	// Teste do codigo da resposta
	w := SetupSaudacao()
	assert.Equal(t, 200, w.Code, "Devem ser iguais")
}

func TestSaudacaoResponseBody(t *testing.T) {
	w := SetupSaudacao()
	body := `{"message":"Eai ga tudo certo"}`
	resp, _ := io.ReadAll(w.Body)
	assert.Equal(t, body, string(resp), "O corpo da resposta tem que estar igual ao requisitado")
}

var Id int

func SetupAlunos() *httptest.ResponseRecorder {
	// Setup da Requisição Saudação
	db.ConectDb()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

var aluno = models.Aluno{
	Nome: "Aluno Teste",
	Cpf:  "12345678910",
	Rg:   "123456789",
}

func CreateAlunoMock() {
	db.Db.Create(&aluno)
	Id = int(aluno.ID)
}

func DeleteAlunoMock() {
	var aluno models.Aluno
	db.Db.Delete(&aluno, Id)
}

func TestAlunosStatusCode(t *testing.T) {
	w := SetupAlunos()
	assert.Equal(t, 200, w.Code, "Devem ser iguais")
}

func TestAlunoAlunoNaLista(t *testing.T) {
	var a []models.Aluno
	CreateAlunoMock()
	defer DeleteAlunoMock()
	existe := false
	w := SetupAlunos()
	resp, _ := io.ReadAll(w.Body)
	json.Unmarshal(resp, &a)
	for _, v := range a {
		if v.Cpf == aluno.Cpf {
			existe = true
		}
	}
	fmt.Println(existe)
	assert.Equal(t, true, existe, "Se existir alguém com o CNPJ requerido o teste está correto")
}
