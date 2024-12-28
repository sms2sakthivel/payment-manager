package model

import "gorm.io/gorm"

type PaymentMode struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey"`
	Mode string `gorm:"unique;not null"`
}

func (pm *PaymentMode) GetAPIResponseObject() *PaymentModeResponse {
	return &PaymentModeResponse{ID: pm.ID, Mode: pm.Mode}
}
