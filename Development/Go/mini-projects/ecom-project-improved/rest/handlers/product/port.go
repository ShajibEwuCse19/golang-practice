package product

import "ecom-project-improved/domain"

// Service defines the contract required by the product HTTP handler.
// This decouples the handler from concrete service/repository implementations.
type Service interface {
	Create(p domain.Product) (*domain.Product, error)
	Get(id uint8) (*domain.Product, error)
	List(page, limit int) ([]*domain.Product, error)
	Delete(id uint8) error
	Update(p domain.Product) (*domain.Product, error)
}