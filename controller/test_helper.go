package controller

import "github.com/bitschain/panicgo/service"

// GetTestController new one controller for test
func GetTestController() *PanicController {
	var c PanicController
	c.Service = service.GetTestService()
	return &c
}
