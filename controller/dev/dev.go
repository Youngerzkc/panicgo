package dev

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	log.Println("ping pong")
	c.String(http.StatusOK, "pong")
}