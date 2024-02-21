package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gogochat/controller"
	"gogochat/tool"
)

func main() {
	engine := gin.Default()

	config := tool.ParseConfig("config")
	if config == nil {
		fmt.Println("config初始化失败")
		return
	}

	registerRouter(engine)

	//初始化ormEngine
	tool.InitDB(config)
	//初始化swagger
	tool.InitSwagger(engine)
	//初始化redis
	tool.InitRedis(config)

	engine.Run(":" + config.App.Port)
}

func registerRouter(engine *gin.Engine) {
	registerHelloController(engine)
	registerUserController(engine)
}
func registerHelloController(engine *gin.Engine) {
	new(controller.HelloController).Router(engine)
}

// 用户模块
func registerUserController(engine *gin.Engine) {
	new(controller.UserController).Router(engine)
}
