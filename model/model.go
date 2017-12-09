package model

import (
	"github.com/jinzhu/gorm"
)

type PanicModel struct {
	DB *gorm.DB
}

func NewModel(db *gorm.DB) *PanicModel {
	var pm PanicModel
	pm.DB = db
	return &pm
}
