package database

import (
	"github.com/bitschain/panicgo/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

const (
	MySQLDialect = "mysql"
)

func NewDB(cfg *config.DBConfig) (*gorm.DB, error) {
	if cfg == nil {
		panic("expect valid db config")
	}

	uri := cfg.ParseURI()
	db, err := gorm.Open(MySQLDialect, uri)
	if err != nil {
		log.Fatalln("failed to open db", uri)
		return db, err
	}
	log.Println("opened database:", cfg.DBName)
	return db, nil
}
