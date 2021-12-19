package entity

type Cart struct {
	ID int64 `json:"id"`
	Items []CartItem `json:"items"`
}

type CartItem struct {
	Product Product `json:"product"`
	Quantity int `json:"quantity"`
}
