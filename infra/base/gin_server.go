package base

import (
	"github.com/gin-gonic/gin"
	"han-xuefeng/gin-scaffold/infra"
)


var ginApplication *gin.Engine

type GinServerStarter struct {
	infra.BaseStarter
}

func Gin() *gin.Engine {
	return initGin()
}

func (g *GinServerStarter) Init(cxt infra.StarterContext) {
	ginApplication = Gin()
}

func (g *GinServerStarter) Setup(cxt infra.StarterContext) {
	ginApplication.GET("/", func(context *gin.Context) {
		context.JSON(
			200,
			gin.H{
				"message": "pong",
			})
	})
}

func (g *GinServerStarter) Start(cxt infra.StarterContext) {
	//把路由打印到控制台
	ginApplication.Run()
}

func (g *GinServerStarter) StartBlocking() bool {
	return true
}

func initGin() *gin.Engine {
	app := gin.New()
	return app
}
