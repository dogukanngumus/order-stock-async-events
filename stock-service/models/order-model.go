package models

type Order struct {
	Id        string `json:"id"`
	ProductId string `json:"productId"`
	Quantity  int    `json:"quantity"`
	CreatedAt string `json:"createdAt"`
}
