package model

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	ID            uint `gorm:"primaryKey"`
	OrderID       uint `gorm:"not null"`
	Amount        uint `gorm:"not null"`
	PaymentModeID uint `gorm:"not null"`

	// Foreign key constraints
	PaymentMode PaymentMode `gorm:"forignKey:PaymentModeID;reference:ID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE;not null"`
}

func (payment *Payment) GetAPIResponseObject() *PaymentResponse {
	return &PaymentResponse{ID: payment.ID, OrderID: payment.OrderID, Amount: payment.Amount, PaymentMode: *payment.PaymentMode.GetAPIResponseObject()}
}
