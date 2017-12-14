package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"path/filepath"
)

// Environment Const
const (
	EnvDevelopment = "development"
	EnvTesting     = "testing"
	EnvDocker      = "docker"
	EnvProduction  = "production"
)

// TomlConfig Global config with toml format
type TomlConfig struct {
	Basic BasicInfo `toml:"basic"`
	DB    DBConfig  `toml:"database"`
}

// BasicInfo Basic config info
type BasicInfo struct {
	Name    string `toml:"name"`
	Version string `toml:"version"`
	ENV     string `toml:"env"`
}

// DBConfig Database config
type DBConfig struct {
	Host      string `toml:"host"`
	Port      uint   `toml:"port"`
	Username  string `toml:"username"`
	Password  string `toml:"password"`
	DBName    string `toml:"dbname"`
	Charset   string `toml:"charset"`
	ParseTime bool   `toml:"parse_time"`
}

// Load load global config
func (c *TomlConfig) Load() error {
	env := os.Getenv("GOENV")
	if env == "" {
		env = EnvDevelopment
	}
	// path := os.Getenv("GOPATH")
	path,_:=os.Getwd()
	// tomlFile := filepath.Join(path, "src/github.com/bitschain/panicgo/config", env+".toml")
	tomlFile:=filepath.Join(path,"/config",env+".toml")
	log.Println("loading config:", tomlFile)
	if _, err := toml.DecodeFile(tomlFile, c); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// ParseURI parse database uri from config
func (dbCfg *DBConfig) ParseURI() string {
	//root:19842895@tcp(localhost:3306)/panicgo?charset=utf8&parseTime=true
	uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t", dbCfg.Username, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.DBName, dbCfg.Charset, dbCfg.ParseTime)
	return uri
}
