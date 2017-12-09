package controller

import "github.com/bitschain/panicgo/service"

// PanicController controller
type PanicController struct {
	Service *service.PanicService
}

// NewController create new controller
func NewController(service *service.PanicService) *PanicController {
	var pc PanicController
	pc.Service = service
	return &pc
}
