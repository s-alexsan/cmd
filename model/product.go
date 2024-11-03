package model

type Product struct {
	ID int `json:"id_product"`
	Name string `json:"product"`
	Price float64 `json:"price"`
}