package dev

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Ping(c *gin.Context) {
	log.Println("ping pong")
	c.String(http.StatusOK, "pong")
}
