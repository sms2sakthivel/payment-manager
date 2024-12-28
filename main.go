package main

import (
	_ "github.com/sms2sakthivel/payment-manager/docs"
	"github.com/sms2sakthivel/payment-manager/payments"
)

func main() {
	// Step 1: Create a New Payment Service Application
	app := payments.NewApp()

	// Step 2: Start Server and Listen on the Port 8001
	app.Listen(":8004")
}
