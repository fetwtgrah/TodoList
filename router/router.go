package router

import (
	"TodoList/controller"

	"github.com/gin-gonic/gin"
)

func StartRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "./static")
	r.StaticFile("/", "./static/index.html")
	v1 := r.Group("/api/v1")
	{
		v1.POST("/todo", controller.CreatTodo)
		v1.GET("/todo", controller.GetAllTodo)
		v1.PUT("/todo/:id", controller.UpdateTodo)
		v1.DELETE("/todo/:id", controller.DeleteTodo)

	}
	return r
}
