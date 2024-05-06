package product

import (
	"errors"
	"gorm.io/gorm"
)

type Repository interface {
	GetAllProducts() ([]Product, error)
	GetProductByID(id string) (*Product, error)
	CreateProduct(product *Product) (*Product, error)
	UpdateProduct(existingProduct *Product, updatedProduct *Product) (*Product, error)
	DeleteProduct(id string) error
}

// ProductRepository เป็น struct ที่ใช้เก็บเมธอดที่ใช้ในการจัดการข้อมูลสินค้าในฐานข้อมูล
type ProductRepository struct {
	db *gorm.DB
}

// NewProductRepository เป็นฟังก์ชันสำหรับสร้าง instance ของ ProductRepository
func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}

// GetAllProducts เรียกข้อมูลสินค้าทั้งหมดจากฐานข้อมูล
func (r *ProductRepository) GetAllProducts() ([]Product, error) {
	var products []Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// GetProductByID ดึงข้อมูลสินค้าจากฐานข้อมูลตาม ID ที่ระบุ
func (r *ProductRepository) GetProductByID(id string) (*Product, error) {
	var product Product
	if err := r.db.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

// CreateProduct สร้างสินค้าใหม่ในฐานข้อมูล
func (r *ProductRepository) CreateProduct(product *Product) (*Product, error) {
	if err := r.db.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

// UpdateProduct อัพเดทข้อมูลสินค้าในฐานข้อมูล
func (r *ProductRepository) UpdateProduct(existingProduct *Product, updatedProduct *Product) (*Product, error) {
	if err := r.db.Model(existingProduct).Updates(updatedProduct).Error; err != nil {
		return nil, err
	}
	return updatedProduct, nil
}

// DeleteProduct ลบข้อมูลสินค้าจากฐานข้อมูล
func (r *ProductRepository) DeleteProduct(id string) error {
	if err := r.db.Delete(&Product{}, id).Error; err != nil {
		return err
	}
	return nil
}
