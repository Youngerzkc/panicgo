package test

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"os"
	"github.com/bitschain/panicgo/server/model"
	"log"
)

func InitTestDB() {
	db, err := gorm.Open("mysql", "root:19842895@tcp(localhost:3306)/panicgo_test?charset=utf8&parseTime=true")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	model.DB = db
	log.Println("init success")
}
