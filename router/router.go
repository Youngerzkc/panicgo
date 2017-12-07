package router

import (
	"github.com/bitschain/panicgo/controller/dev"
	"github.com/bitschain/panicgo/controller/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(e *gin.Engine) {
	// Client group: router
	devRouter := e.Group("/dev")
	{
		devRouter.GET("/ping", dev.Ping)
	}

	userRouter := e.Group("/user")
	{
		userRouter.POST("/signup", user.Signup)
		userRouter.POST("/signin", user.Signin)
		userRouter.POST("/signout", user.Signout)
		userRouter.GET("/:id", user.PublicInfo)
		userRouter.PUT("/:id", user.UpdateInfo)
	}
}
