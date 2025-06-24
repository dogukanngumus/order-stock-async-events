package service

import (
	"fmt"
	"stock-service/models"
)

var stockData = map[string]int{
	"product-1": 100,
	"product-2": 50,
}

func ProcessOrder(order models.Order) {
	stockData[order.ProductId] -= order.Quantity
	fmt.Printf("✅ Sipariş işlendi: Ürün %s, Miktar %d → Kalan stok: %d\n",
		order.ProductId, order.Quantity, stockData[order.ProductId])
}
