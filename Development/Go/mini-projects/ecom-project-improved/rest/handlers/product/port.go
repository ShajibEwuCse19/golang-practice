package product

import "ecom-project-improved/domain"

type Service interface {
	Create(p domain.Product) (*domain.Product, error)
	Get(id uint8) (*domain.Product, error)
	List() ([]*domain.Product, error)
	Delete(id uint8) error
	Update(p domain.Product) (*domain.Product, error)
}