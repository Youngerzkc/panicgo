package database

import (
	"github.com/bitschain/panicgo/config"
	"github.com/jinzhu/gorm"
	"os"
)

func GetTestDB() *gorm.DB {
	os.Setenv("GOENV", config.EnvTesting)
	var cfg config.TomlConfig
	cfg.Load()
	db, _ := NewDB(&cfg.DB)
	return db
}
