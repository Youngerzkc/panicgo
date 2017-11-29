package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/leonzhao/panicgo/router"
)

func main() {
	app := gin.Default()

	router.InitRoutes(app)

	app.Run(":8080")

	log.Println("server is listening :8080")

}

