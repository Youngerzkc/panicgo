package service

import (
	"github.com/bitschain/panicgo/model"
)

type PanicService struct {
	Model *model.PanicModel
}

func NewService(pm *model.PanicModel) *PanicService {
	var service PanicService
	service.Model = pm
	return &service
}
