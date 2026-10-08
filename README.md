# TodoList

一个基于 Go + Gin + GORM + PostgreSQL 实现的简单 Todo List 项目。

这是一个用于学习 Go Web 后端以及前后端联调的练习项目。

## 🛠️ 技术栈

### 后端

- Go
- Gin
- GORM
- PostgreSQL
- zap

### 前端

- HTML
- CSS
- JavaScript
- Fetch API

### 开发工具

- Git
- GoLand / VS Code
- Postman / 浏览器开发者工具

---

## 📁 项目结构

```text
TodoList/
│
├── configs/
│   └── database.go          # 数据库连接配置
│
├── controller/
│   └── todo.go              # Todo 业务处理
│
├── model/
│   └── todo.go              # Todo 数据模型
│
├── router/
│   └── router.go            # 路由配置
│
├── static/
│   └── index.html           # 前端页面
│
├── go.mod
├── go.sum
│
└── main.go                  # 项目入口
