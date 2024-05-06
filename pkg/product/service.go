package product

import (
	"errors"
)

// Service เป็น struct ที่ใช้เก็บ business logic ของสินค้า
type Service struct {
	repository Repository
}

// NewService เป็นฟังก์ชันสำหรับสร้าง instance ของ Service
func NewService(repository Repository) *Service {
	return &Service{repository}
}

// GetAllProducts ใช้สำหรับดึงข้อมูลสินค้าทั้งหมด
func (s *Service) GetAllProducts() ([]Product, error) {
	// เรียกใช้ method ใน repository เพื่อดึงข้อมูลสินค้าทั้งหมด
	return s.repository.GetAllProducts()
}

// GetProductByID ใช้สำหรับดึงข้อมูลสินค้าตาม ID
func (s *Service) GetProductByID(id string) (*Product, error) {
	// เรียกใช้ method ใน repository เพื่อดึงข้อมูลสินค้าตาม ID
	return s.repository.GetProductByID(id)
}

// CreateProduct ใช้สำหรับสร้างสินค้าใหม่
func (s *Service) CreateProduct(product *Product) (*Product, error) {
	// เรียกใช้ method ใน repository เพื่อสร้างสินค้าใหม่
	return s.repository.CreateProduct(product)
}

// UpdateProduct ใช้สำหรับอัพเดทข้อมูลสินค้า
func (s *Service) UpdateProduct(id string, updatedProduct *Product) (*Product, error) {
	// ตรวจสอบว่าสินค้าที่ต้องการอัพเดทมีอยู่หรือไม่
	existingProduct, err := s.repository.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	if existingProduct == nil {
		return nil, errors.New("product not found")
	}

	// เรียกใช้ method ใน repository เพื่ออัพเดทสินค้า
	return s.repository.UpdateProduct(existingProduct, updatedProduct)
}

// DeleteProduct ใช้สำหรับลบข้อมูลสินค้า
func (s *Service) DeleteProduct(id string) error {
	// เรียกใช้ method ใน repository เพื่อลบข้อมูลสินค้า
	return s.repository.DeleteProduct(id)
}
