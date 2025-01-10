package models

type CartItem struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

var Carts = map[uint][]CartItem{} // Map user_id ke list keranjang belanja
