package model

type PaymentCreateRequest struct {
	OrderID       uint `json:"order_id"`
	Amount        uint `json:"amount"`
	PaymentModeID uint `json:"payment_mode_id"`
}

type PaymentUpdateRequest struct {
	ID            uint `json:"payment_id"`
	OrderID       uint `json:"order_id"`
	Amount        uint `json:"amount"`
	PaymentModeID uint `json:"payment_mode_id"`
}

func (pcr *PaymentCreateRequest) GetDBObject() *Payment {
	return &Payment{OrderID: pcr.OrderID, Amount: pcr.Amount, PaymentModeID: pcr.PaymentModeID}
}

func (pur *PaymentUpdateRequest) GetDBObject() *Payment {
	return &Payment{ID: pur.ID, OrderID: pur.OrderID, Amount: pur.Amount, PaymentModeID: pur.PaymentModeID}
}

type PaymentResponse struct {
	ID          uint                `json:"payment_id"`
	OrderID     uint                `json:"order_id"`
	Amount      uint                `json:"amount"`
	PaymentMode PaymentModeResponse `json:"payment_mode"`
}
