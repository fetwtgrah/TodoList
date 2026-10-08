package controller

import (
	"TodoList/configs"
	"TodoList/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UpdateTodoRequest struct {
	Content string `json:"content"`
	Status  string `json:"status" binding:"omitempty,oneof=done pending"`
}
type TodoController struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewTodoController(db *gorm.DB, logger *zap.Logger) *TodoController {
	return &TodoController{
		db:     db,
		logger: logger,
	}
}

func (t *TodoController) CreateTodo(c *gin.Context) {
	var todo model.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		t.logger.Warn("创建todo命令无效",
			zap.Error(err))

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "无效输入",
		})
		return
	}
	result := t.db.Create(&todo)
	if result.Error != nil {
		t.logger.Warn("创建todo失败",
			zap.Error(result.Error))
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": result.Error.Error(),
		})
		return
	}
	t.logger.Info("创建todo成功",
		zap.Uint("todoID", todo.ID))
	c.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"data": todo,
	})
}

func (t *TodoController) GetAllTodo(c *gin.Context) {
	var todos []model.Todo
	err := t.db.Find(&todos).Error
	if err != nil {
		t.logger.Warn("数据获取出错",
			zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "数据获取出错",
		})
		return
	}
	t.logger.Info("成功获取todo list")
	c.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"data": todos,
	})
}

func (t *TodoController) UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var req UpdateTodoRequest
	var dbTodo model.Todo
	if err := c.ShouldBindJSON(&req); err != nil {
		t.logger.Warn("参数设置失败",
			zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "无效数据",
		})
		return
	}
	if err := t.db.First(&dbTodo, id).Error; err != nil {
		t.logger.Warn("todo查找失败",
			zap.Error(err))
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

	if err := t.db.Save(&dbTodo).Error; err != nil {
		t.logger.Warn("数据更新失败",
			zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	t.logger.Info("数据更新成功",
		zap.Uint("todoID", dbTodo.ID))
	c.JSON(http.StatusOK, gin.H{
		"msg":  "ok,更新成功",
		"data": dbTodo,
	})
}

func (t *TodoController) DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	var todo model.Todo
	if err := t.db.First(&todo, id).Error; err != nil {
		t.logger.Warn("todo查找失败",
			zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	if err := configs.Db.Delete(&todo).Error; err != nil {
		t.logger.Warn("todo删除失败",
			zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	t.logger.Info("删除成功",
		zap.String("todoID", id))
	c.JSON(http.StatusOK, gin.H{
		"msg": "删除成功",
	})
}
