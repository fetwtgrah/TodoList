package router

import (
	"TodoList/controller"

	"github.com/gin-gonic/gin"
)

func StartRouter(t *controller.TodoController) *gin.Engine {
	r := gin.Default()
	r.Static("/static", "./static")
	r.StaticFile("/", "./static/index.html")
	v1 := r.Group("/api/v1")
	{
		v1.POST("/todo", t.CreateTodo)
		v1.GET("/todo", t.GetAllTodo)
		v1.PUT("/todo/:id", t.UpdateTodo)
		v1.DELETE("/todo/:id", t.DeleteTodo)

	}
	return r
}
