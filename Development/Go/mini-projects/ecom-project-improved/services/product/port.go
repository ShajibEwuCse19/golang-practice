package services

import (
	"ecom-project-improved/domain"
	productHandler "ecom-project-improved/rest/handlers/product"
)

type Service interface {
	productHandler.Service
}

// independent ProductRepo ports [contract] -> repository implements this interface
type ProductRepo interface {
	Create(p domain.Product) (*domain.Product, error)
	Get(id uint8) (*domain.Product, error)
	List(page, limit int) ([]*domain.Product, error)
	Delete(id uint8) error
	Update(p domain.Product) (*domain.Product, error)
}
