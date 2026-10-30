package product

import "ecom-project-improved/domain"

// Introduced contract list for product handler.
// This is solve the direct dependent with product repo. 
type Service interface {
	Create(p domain.Product) (*domain.Product, error)
	Get(id uint8) (*domain.Product, error)
	List(page, limit int) ([]*domain.Product, error)
	Delete(id uint8) error
	Update(p domain.Product) (*domain.Product, error)
}