package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/llRedXD/Go6/controllers"
)

func Handler() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/alunos/:id", controllers.ExibeAluno)
	r.GET("/alunos/cpf/:cpf", controllers.SearchAlunoCpf)
	r.POST("/criaAlunos", controllers.CriarAluno)
	r.DELETE("/deleteAlunos/:id", controllers.DeleteAluno)
	r.PUT("/updateAlunos/:id", controllers.UpdateAluno)
	r.GET("/:nome", controllers.Saudacoes)
	r.Run()
}
