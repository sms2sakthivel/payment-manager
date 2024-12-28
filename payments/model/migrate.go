package model

import (
	"gorm.io/gorm"
)

func Automigrate(DB *gorm.DB) error {
	return DB.AutoMigrate(&Payment{})
}
