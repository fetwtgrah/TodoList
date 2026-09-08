package main

import (
	"TodoList/configs"
	"TodoList/model"
	"TodoList/router"
	"fmt"
)

func main() {
	if err := configs.InitConfig(); err != nil {
		fmt.Sprintf("系统文件配置出错,%s", err)
	}
	if err := configs.InitDb(); err != nil {
		fmt.Sprintf("数据库初始化失败,%v", err)
	}
	err := configs.Db.AutoMigrate(&model.Todo{})
	if err != nil {
		fmt.Println("数据表创建失败")
	}
	r := router.StartRouter()
	r.Run(":8080")
}
