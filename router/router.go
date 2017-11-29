package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leonzhao/panicgo/controller/dev"
)

func InitRoutes(r *gin.Engine) {
	// Client group: router
	devRouter := r.Group("/dev")
	{
		devRouter.GET("/ping", dev.Ping)
	}
}