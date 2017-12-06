package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/bitschain/panicgo/server/model"
	"github.com/bitschain/panicgo/server/router"
)

func init() {
	db, err := gorm.Open("mysql", "root:19842895@tcp(localhost:3306)/panicgo?charset=utf8&parseTime=true")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	model.DB = db
	log.Println("init success")
}

func main() {
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
