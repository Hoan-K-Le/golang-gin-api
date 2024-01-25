package models

type Product struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Quantity int64 `json:"quantity"`
	Category string `json:"category"`
	ImageUrl string `json:"imgurl"`
}
