package service

import (
	"github.com/bitschain/panicgo/database"
	"github.com/bitschain/panicgo/model"
)

func GetTestService() *PanicService {
	var pm model.PanicModel
	pm.DB = database.GetTestDB()

	var s PanicService
	s.Model = &pm
	return &s
}
