package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Ping ping pong test
func (pc *PanicController) Ping(c *gin.Context) {
	log.Println("ping pong")
	c.String(http.StatusOK, "pong")
}
