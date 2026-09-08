package controller

import (
	"TodoList/configs"
	"TodoList/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateTodoRequest struct {
	Content string `json:"content"`
	Status  string `json:"status" binding:"omitempty,oneof=done pending"`
}

func CreatTodo(c *gin.Context) {
	var todo model.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "无效输入",
		})
		return
	}
	result := configs.Db.Create(&todo)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"data": todo,
	})
}

func GetAllTodo(c *gin.Context) {
	var todos []model.Todo
	err := configs.Db.Find(&todos).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "数据获取出错",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"data": todos,
	})
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var req UpdateTodoRequest
	var dbTodo model.Todo

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "无效数据",
		})
		return
	}

	if err := configs.Db.First(&dbTodo, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	if req.Content != "" {
		dbTodo.Name = req.Content
	}
	if req.Status != "" {
		dbTodo.Status = req.Status
	}

	if err := configs.Db.Save(&dbTodo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "ok,更新成功",
		"data": dbTodo,
	})
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	var todo model.Todo
	if err := configs.Db.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	if err := configs.Db.Delete(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "删除成功",
	})
}
