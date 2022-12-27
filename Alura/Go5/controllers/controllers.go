package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/llRedXD/go5/db"
	"github.com/llRedXD/go5/models"
)

func ExibeTodosAlunos(c *gin.Context) {
	a := []models.Aluno{}
	db.Db.Find(&a)
	c.JSON(200, a)
}

func ExibeAluno(c *gin.Context) {
	a := models.Aluno{}
	id := c.Params.ByName("id")
	db.Db.Find(&a, id)
	if a.ID == 0 {
		c.JSON(400, gin.H{"Not Found": "Aluno não encontrado"})
		return
	}
	c.JSON(200, a)
}

func CriarAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(400, gin.H{"error": "Dados inválidos"})
		return
	}
	db.Db.Create(&aluno)
	c.JSON(200, aluno)
}

func DeleteAluno(c *gin.Context) {
	a := models.Aluno{}
	id := c.Params.ByName("id")
	db.Db.Find(&a, id)
	fmt.Println(a)
	if a.ID == 0 {
		c.JSON(400, gin.H{"Not Found": "Aluno não encontrado"})
		return
	}
	db.Db.Delete(&a, id)
	c.JSON(200, a)
}

func UpdateAluno(c *gin.Context) {
	a := models.Aluno{}
	id := c.Params.ByName("id")
	db.Db.Find(&a, id)
	if a.ID == 0 {
		c.JSON(400, gin.H{"Not Found": "Aluno não encontrado"})
		return
	}
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(400, gin.H{"error": "Dados inválidos"})
		return
	}
	db.Db.Model(&a).Updates(a)
	c.JSON(200, a)
}

func SearchAlunoCpf(c *gin.Context) {
	a := models.Aluno{}
	cpf := c.Params.ByName("cpf")
	db.Db.Where("cpf = ?", cpf).Find(&a)
	if a.ID == 0 {
		c.JSON(400, gin.H{"Not Found": "Aluno não encontrado"})
		return
	}
	c.JSON(200, a)
}

func Saudacoes(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"message": "Eai " + nome + " tudo certo",
	})
}
