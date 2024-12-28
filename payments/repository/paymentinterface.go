package repository

import (
	"github.com/sms2sakthivel/payment-manager/payments/model"
)

type PaymentRepository interface {
	GetAllPayments() ([]model.Payment, error)
	GetPaymentByID(id uint) (*model.Payment, error)
	CreatePayment(Payment *model.Payment) error
	UpdatePayment(Payment *model.Payment) error
	DeletePayment(id uint) error
}
