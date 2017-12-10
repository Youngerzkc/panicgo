package database

import (
	"github.com/bitschain/panicgo/config"
	"github.com/jinzhu/gorm"
)

func GetTestDB() *gorm.DB {
	var cfg config.TomlConfig
	cfg.Load()
	db, _ := NewDB(&cfg.DB)
	return db
}
