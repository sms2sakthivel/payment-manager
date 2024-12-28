package payments

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sms2sakthivel/payment-manager/payments/database"
	"github.com/sms2sakthivel/payment-manager/payments/model"
	"github.com/sms2sakthivel/payment-manager/payments/repository"
	"github.com/sms2sakthivel/payment-manager/payments/routes"
	"github.com/sms2sakthivel/payment-manager/payments/service"
)

func NewApp() *fiber.App {
	// Step 1: Connect to the database
	database.Connect()

	// Step 2: Auto-migrate Payment schema
	database.DB.AutoMigrate(&model.Payment{})

	// Step 3: Initialize repository, service, and app
	repo := &repository.GormPaymentRepository{DB: database.DB}
	service := &service.PaymentService{Repo: repo}
	app := fiber.New()

	// Step 4: Enable Logger middleware with timing
	app.Use(logger.New(logger.Config{
		Format: "${time} - ${latency} - ${status} - ${method} ${path}\n",
	}))

	// Step 5: Register routes
	routes.RegisterRoutes(app, service)
	return app
}
