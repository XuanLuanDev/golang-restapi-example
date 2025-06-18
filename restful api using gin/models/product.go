package models

type Product struct {
	ID    int    `json:"id" gorm:"column:id;primaryKey;autoIncrement"`  
	Name  string `json:"name" gorm:"column:name"`
	Price int    `json:"price" gorm:"column:price"`
	Code  string `json:"code" gorm:"column:code"`
	Desc  string `json:"description" gorm:"column:description"`
	Amount int    `json:"amount" gorm:"column:amount"`
	Status string `json:"status" gorm:"column:statud"`
}