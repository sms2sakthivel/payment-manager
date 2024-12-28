package model

type OrderResponse struct {
	ID     uint         `json:"order_id"`
	CartID uint         `json:"cart_id"`
	Cart   CartResponse `json:"cart,omitempty"`
}

type CartResponse struct {
	ID        uint               `json:"cart_id"`
	CartItems []CartItemResponse `json:"cart_items"`
	UserID    uint               `json:"user_id"`
	CartValue uint               `json:"cart_value"`
}

type CartItemResponse struct {
	ID        uint `json:"cart_item_id"`
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
	Discount  int  `json:"discount"`
}
