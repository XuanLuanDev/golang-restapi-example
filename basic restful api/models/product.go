package models

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Code  string `json:"code"`
	Desc  string `json:"description"`
	Amount int    `json:"amount"`
	Status string `json:"status"`
}