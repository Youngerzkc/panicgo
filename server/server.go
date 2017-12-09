package server

import (
	"github.com/bitschain/panicgo/config"
	"github.com/bitschain/panicgo/controller"
	"github.com/bitschain/panicgo/database"
	"github.com/bitschain/panicgo/model"
	"github.com/bitschain/panicgo/service"
	"github.com/jinzhu/gorm"
)

type Context struct {
	DB         *gorm.DB
	Model      *model.PanicModel
	Service    *service.PanicService
	Controller *controller.PanicController
}

type Panicgo struct {
	Context *Context
	Status  string
}

func NewServer(cfg *config.TomlConfig) *Panicgo {
	var c Context
	db, err := database.NewDB(&cfg.DB)
	if err != nil {
		panic(err)
	}
	c.DB = db
	c.Model = model.NewModel(c.DB)
	c.Service = service.NewService(c.Model)
	c.Controller = controller.NewController(c.Service)

	var pg Panicgo
	pg.Context = &c
	return &pg
}

func (c *Context) GetService() {

}
