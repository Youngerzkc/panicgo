package main

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/bitschain/panicgo/model"
	"github.com/bitschain/panicgo/router"
)

func init() {
}

func main() {
	model.OpenDatabase()
	defer model.CloseDatabase()

	routes := gin.Default()

	router.InitRoutes(routes)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ... ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exit")
}
