package model

// Product represents a product entity in the system.
type Product struct {
	ID   	int 		`gorm:"primary_key" json:"id"` // ID of the product (primary key)
	Name 	string		`json:"name"` // Name of the product
	Price 	float32		`json:"price"` // Price of the product
}