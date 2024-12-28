package service

import (
	"github.com/gofiber/fiber/v3"
	"github.com/sms2sakthivel/payment-manager/payments/model"
	"github.com/sms2sakthivel/payment-manager/payments/repository"
)

type PaymentService struct {
	Repo repository.PaymentRepository
}

func (s *PaymentService) GetPayments() ([]model.PaymentResponse, error) {
	dbPayments, err := s.Repo.GetAllPayments()
	if err != nil {
		return nil, err
	}
	var payments []model.PaymentResponse = []model.PaymentResponse{}
	for _, payment := range dbPayments {
		payments = append(payments, *payment.GetAPIResponseObject())
	}
	return payments, nil
}

func (s *PaymentService) GetPayment(id uint) (*model.PaymentResponse, error) {
	payment, err := s.Repo.GetPaymentByID(id)
	if err != nil {
		return nil, err
	}
	return payment.GetAPIResponseObject(), err
}

func (s *PaymentService) CreatePayment(paymentReq *model.PaymentCreateRequest) (*model.PaymentResponse, *fiber.Error) {
	// Step 1: Check if the OrderID is Valid
	order, err := GetOrderByID(paymentReq.OrderID)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// Step 2: Check if the order cart value is equal to the payment amount
	if order.Cart.CartValue != paymentReq.Amount {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Payment & Cart Value Amount Mismatch")
	}
	payment := paymentReq.GetDBObject()
	er := s.Repo.CreatePayment(payment)
	if er != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, er.Error())
	}
	return payment.GetAPIResponseObject(), nil
}

func (s *PaymentService) UpdatePayment(paymentReq *model.PaymentUpdateRequest) (*model.PaymentResponse, *fiber.Error) {
	// Step 1: Check if the OrderID is Valid
	order, err := GetOrderByID(paymentReq.OrderID)
	if err != nil {
		return nil, err
	}
	// Step 2: Check if the order cart value is equal to the payment amount
	if order.Cart.CartValue != paymentReq.Amount {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Payment & Cart Value Amount Mismatch")
	}
	payment := paymentReq.GetDBObject()
	er := s.Repo.UpdatePayment(payment)
	return payment.GetAPIResponseObject(), fiber.NewError(fiber.StatusInternalServerError, er.Error())
}

func (s *PaymentService) DeletePayment(id uint) error {
	return s.Repo.DeletePayment(id)
}
