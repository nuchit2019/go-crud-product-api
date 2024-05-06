package product

type Product struct {
	ID    int     `gorm:"primary_key" json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
