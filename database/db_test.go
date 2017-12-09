package database

import (
	"github.com/bitschain/panicgo/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDB(t *testing.T) {

	cfg := config.DBConfig{
		Host:      "localhost",
		Port:      3306,
		Username:  "root",
		Password:  "123456",
		DBName:    "panicgo_test",
		Charset:   "utf8",
		ParseTime: true,
	}
	db, err := NewDB(&cfg)
	assert.Nil(t, err)
	assert.NotEqual(t, nil, db)
}
