package service

import (
	"fmt"
	"stock-service/models"
)

func ProcessOrder(order models.Order) {
	fmt.Printf("✅ Sipariş işlendi: Ürün %s\n", order.ProductName, order.Quantity)
}
