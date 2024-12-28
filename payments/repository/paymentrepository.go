package repository

import (
	"github.com/sms2sakthivel/payment-manager/payments/model"
	"gorm.io/gorm"
)

type GormPaymentRepository struct {
	DB *gorm.DB
}

func (repo *GormPaymentRepository) GetAllPayments() ([]model.Payment, error) {
	var payments []model.Payment
	err := repo.DB.Preload("PaymentMode").Find(&payments).Error
	return payments, err
}

func (repo *GormPaymentRepository) GetPaymentByID(id uint) (*model.Payment, error) {
	var payment model.Payment
	err := repo.DB.Preload("PaymentMode").First(&payment, id).Error
	return &payment, err
}

func (repo *GormPaymentRepository) CreatePayment(payment *model.Payment) error {
	var paymentMode model.PaymentMode
	err := repo.DB.First(&paymentMode, payment.PaymentModeID).Error
	if err != nil {
		return err
	}
	payment.PaymentMode = paymentMode
	return repo.DB.Create(payment).Error
}

func (repo *GormPaymentRepository) UpdatePayment(payment *model.Payment) error {
	return repo.DB.Save(payment).Error
}

func (repo *GormPaymentRepository) DeletePayment(id uint) error {
	return repo.DB.Delete(&model.Payment{}, id).Error
}
