package server

import (
	"github.com/bitschain/panicgo/config"
	"testing"
	//"github.com/stretchr/testify/assert"
	"os"
)

func TestNewServer(t *testing.T) {
	os.Setenv("GOENV", config.EnvTesting)
	var cfg config.TomlConfig
	cfg.Load()

	s := NewServer(&cfg)
	t.Log(s.Status)

	//assert.NotNil(t, server)

}
