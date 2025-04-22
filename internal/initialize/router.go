package initialize

import (
	"mtb_web/global"
	"mtb_web/internal/middlewares"
	"mtb_web/internal/router"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	//middleware
	r.Use(middlewares.LoggerMiddle())        //logging
	r.Use(middlewares.CorsMiddleware())      //cross
	r.Use(middlewares.RateLimitMiddleware()) //limiter global
	managerRouter := router.RouterGroupApp.Manager
	userRouter := router.RouterGroupApp.User

	mainGroup := r.Group("/api/v1")
	{
		mainGroup.GET("/checkStatus")
	}
	{
		userRouter.InitProductSizeRouter(mainGroup)
		userRouter.InitProductColorRouter(mainGroup)
	}
	{
		managerRouter.InitProductSizeRouter(mainGroup)
		managerRouter.InitProductColorRouter(mainGroup)
	}
	return r
}
