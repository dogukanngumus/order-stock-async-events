package models

type Order struct {
	Id          int    `json:"id"`
	ProductName string `json:"productName"`
	Quantity    int    `json:"quantity"`
	CreatedAt   string `json:"createdAt"`
}
