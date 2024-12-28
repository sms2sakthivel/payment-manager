package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/client"
	"github.com/sms2sakthivel/payment-manager/payments/model"
)

func GetOrderByID(id uint) (*model.OrderResponse, *fiber.Error) {
	cli := client.New()
	cli.SetTimeout(10 * time.Second)

	// Send a GET request
	resp, err := cli.Get(fmt.Sprintf("http://127.0.0.1:8003/orders/%d", id))
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	fmt.Printf("Status: %d\n", resp.StatusCode())
	fmt.Printf("Body: %s\n", string(resp.Body()))
	if resp.StatusCode() != 200 {
		return nil, fiber.NewError(resp.StatusCode(), "Order Not Found")
	}
	var orderResponse model.OrderResponse
	err = json.Unmarshal(resp.Body(), &orderResponse)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return &orderResponse, nil
}
