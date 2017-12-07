package model

import (
	"github.com/bitschain/panicgo/config"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

func OpenDatabase() {
	c, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	connStr := config.ParseDBUrl(c)

	log.Println("connecting database.")
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
		panic("failed to connect database")
	}
	DB = db
}

func CloseDatabase() {
	log.Println("close database.")
	defer DB.Close()
}
