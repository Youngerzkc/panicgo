package router

import (
	"github.com/bitschain/panicgo/server"
	"github.com/gin-gonic/gin"
)

func InitRoutes(e *gin.Engine, c *server.Context) {
	// Client group: router
	devRouter := e.Group("/dev")
	{
		devRouter.GET("/ping", c.Controller.Ping)
	}

	//userRouter := e.Group("/user")
	//{
	//	userRouter.POST("/signup", c.Controller.Signup)
	//	userRouter.POST("/signin", c.Controller.Signin)
	//	userRouter.POST("/signout", c.Controller.Signout)
	//	userRouter.GET("/:id", c.Controller.PublicInfo)
	//	userRouter.PUT("/:id", c.Controller.UpdateInfo)
	//}
}
