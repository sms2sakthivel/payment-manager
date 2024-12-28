package routes

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/sms2sakthivel/payment-manager/payments/model"
	"github.com/sms2sakthivel/payment-manager/payments/service"
)

func RegisterRoutes(app *fiber.App, service *service.PaymentService) {
	app.Get("/", PaymentServiceInfo)
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/payments", func(c *fiber.Ctx) error { return GetAllPayments(c, service) })
	app.Get("/payments/:id", func(c *fiber.Ctx) error { return GetPaymentByID(c, service) })
	app.Post("/payments", func(c *fiber.Ctx) error { return CreatePayment(c, service) })
	app.Put("/payments/:id", func(c *fiber.Ctx) error { return UpdatePayment(c, service) })
	app.Delete("/payments/:id", func(c *fiber.Ctx) error { return DeletePayment(c, service) })
}

// PaymentServiceInfo returns information about the Payment Service
//
// @Summary      Payment Service Info
// @Description  Returns basic information about the Payment Service
// @Tags         General
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       / [get]
func PaymentServiceInfo(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Payment Service"})
}

// GetAllPayments retrieves all payments
//
// @Summary      Get All Payments
// @Description  Retrieve a list of all payments
// @Tags         Payments
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.PaymentResponse
// @Failure      500  {object}  fiber.Error
// @Router       /payments [get]
func GetAllPayments(c *fiber.Ctx, service *service.PaymentService) error {
	payments, err := service.GetPayments()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(payments)
}

// GetPaymentByID retrieves a payment by their ID
//
// @Summary      Get Payment by ID
// @Description  Retrieve a payment by their ID
// @Tags         Payments
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Payment ID"
// @Success      200  {object}  model.PaymentResponse
// @Failure      404  {object}  fiber.Error
// @Failure      500  {object}  fiber.Error
// @Router       /payments/{id} [get]
func GetPaymentByID(c *fiber.Ctx, service *service.PaymentService) error {
	id, _ := strconv.Atoi(c.Params("id"))
	payment, err := service.GetPayment(uint(id))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Payment not found")
	}
	return c.JSON(payment)
}

// CreatePayment adds a new payment
//
// @Summary      Create a New Payment
// @Description  Add a new payment to the system
// @Tags         Payments
// @Accept       json
// @Produce      json
// @Param        payment  body      model.PaymentCreateRequest  true  "Payment details"
// @Success      201   {object}  model.PaymentResponse
// @Failure      400   {object}  fiber.Error
// @Failure      500   {object}  fiber.Error
// @Router       /payments [post]
func CreatePayment(c *fiber.Ctx, service *service.PaymentService) error {
	var paymentReq model.PaymentCreateRequest
	if err := c.BodyParser(&paymentReq); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}
	payment, err := service.CreatePayment(&paymentReq)
	if err != nil {
		return c.Status(err.Code).JSON(err.Error())
		// return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(payment)
}

// UpdatePayment modifies details of an existing payment
//
// @Summary      Update an Existing Payment
// @Description  Modify details of an existing payment
// @Tags         Payments
// @Accept       json
// @Produce      json
// @Param        id    path      int         true  "Payment ID"
// @Param        payment  body      model.PaymentUpdateRequest  true  "Updated payment details"
// @Success      200   {object}  model.PaymentResponse
// @Failure      400   {object}  fiber.Error
// @Failure      404   {object}  fiber.Error
// @Failure      500   {object}  fiber.Error
// @Router       /payments/{id} [put]
func UpdatePayment(c *fiber.Ctx, service *service.PaymentService) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var paymentReq model.PaymentUpdateRequest
	if err := c.BodyParser(&paymentReq); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}
	paymentReq.ID = uint(id)
	payment, err := service.UpdatePayment(&paymentReq)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(payment)
}

// DeletePayment removes a payment by their ID
//
// @Summary      Delete a Payment
// @Description  Remove a payment by their ID
// @Tags         Payments
// @Accept       json
// @Produce      json
// @Param        id   path  int  true  "Payment ID"
// @Success      204
// @Failure      500  {object}  fiber.Error
// @Router       /payments/{id} [delete]
func DeletePayment(c *fiber.Ctx, service *service.PaymentService) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := service.DeletePayment(uint(id)); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}
