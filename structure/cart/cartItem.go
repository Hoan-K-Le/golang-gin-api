package cart

type CartItem struct {
	ProductID int64 `json:"productId"`
	Quantity int `json:"quantity"`
	Category string `json:"category"`
}