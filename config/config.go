package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"path/filepath"
)

const (
	envDevelopment = "development"
	envTesting     = "testing"
	envDocker      = "docker"
	envProduction  = "production"
)

type Config struct {
	Basic basicInfo `toml:"basic"`
	DB    database  `toml:"database"`
}

type basicInfo struct {
	Name    string `toml:"name"`
	Version string `toml:"version"`
	ENV     string `toml:"env"`
}

type database struct {
	Host      string `toml:"host"`
	Port      uint   `toml:"port"`
	Username  string `toml:"username"`
	Password  string `toml:"password"`
	DBName    string `toml:"dbname"`
	Charset   string `toml:"charset"`
	ParseTime bool   `toml:"parse_time"`
}

func GetConfig() (Config, error) {
	env := os.Getenv("GOENV")
	if env == "" {
		env = envDevelopment
	}
	path := os.Getenv("GOPATH")

	tomlFile := filepath.Join(path, "src/github.com/bitschain/panicgo/config", env+".toml")
	log.Println(tomlFile)

	var config Config
	if _, err := toml.DecodeFile(tomlFile, &config); err != nil {
		log.Fatal(err)
		return config, err
	}

	return config, nil
}

func ParseDBUrl(c Config) string {
	//root:19842895@tcp(localhost:3306)/panicgo?charset=utf8&parseTime=true
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t", c.DB.Username, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.DBName, c.DB.Charset, c.DB.ParseTime)
	return connStr
}
