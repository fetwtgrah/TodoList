package main

import (
	"TodoList/configs"
	"fmt"
)

func main() {
	if err := configs.InitConfig(); err != nil {
		fmt.Sprintf("系统文件配置出错,%s", err)
		return
	}
	if err := configs.InitDb(); err != nil {
		fmt.Sprintf("数据库初始化失败,%v", err)
		return
	}
	fmt.Println("测试通过")

}
