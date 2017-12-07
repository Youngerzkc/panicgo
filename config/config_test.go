package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetConfigDevelopmentEnv(t *testing.T) {
	c, err := GetConfig()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "panicgo", c.Basic.Name)
	assert.Equal(t, "development", c.Basic.ENV)
}

func TestGetConfigProductionEnv(t *testing.T) {
	os.Setenv("GOENV", "production")
	c, err := GetConfig()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "panicgo", c.Basic.Name)
	assert.Equal(t, "production", c.Basic.ENV)
}

func TestParseDBUrl(t *testing.T) {
	c, err := GetConfig()
	if err != nil {
		t.Fatal(err)
	}
	connStr := ParseDBUrl(c)
	assert.Equal(t, "root:123456@tcp(localhost:3306)/panicgo?charset=utf8&parseTime=true", connStr)
}
