package main

import (
	"TodoList/configs"
	"TodoList/controller"
	"TodoList/logger"
	"TodoList/model"
	"TodoList/router"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func main() {
	log := logger.InitLogger()
	defer func(log *zap.Logger) {
		_ = log.Sync()
	}(log)
	if err := configs.InitConfig(); err != nil {
		log.Fatal("配置文件读取出错", zap.Error(err))
	}

	if err := configs.InitDb(); err != nil {
		log.Fatal("数据库初始化失败", zap.Error(err))
	}

	if err := configs.Db.AutoMigrate(&model.Todo{}); err != nil {
		log.Fatal("数据表创建失败", zap.Error(err))
	}
	t := controller.NewTodoController(configs.Db, log)

	r := router.StartRouter(t)

	s := &http.Server{
		Addr:           configs.Conf.Server.Port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Error("服务器启动失败", zap.Error(err))
	}
}
