package models

type Product struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var Products = []Product{
	{ID: 1, Name: "Product A", Price: 10000},
	{ID: 2, Name: "Product B", Price: 20000},
}
