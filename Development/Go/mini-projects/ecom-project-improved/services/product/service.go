package services

import "ecom-project-improved/domain"

type service struct {
	productRepo ProductRepo
}

// NewService initializes a new product service with its repository dependency.
// It returns the Service interface defined in services/product/port.go.
func NewService(productRepo ProductRepo) Service {
	return &service{
		productRepo: productRepo,
	}
}

// The service struct implements all methods of the Service interface.
func (svc *service) Create(p domain.Product) (*domain.Product, error) {
	prd, err := svc.productRepo.Create(p)
	if err != nil {
		return nil, err
	}
	return prd, nil
}
func (svc *service) Delete(id uint8) error {
	err := svc.productRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
func (svc *service) Get(id uint8) (*domain.Product, error) {
	prd, err := svc.productRepo.Get(id)
	if err != nil {
		return nil, err
	}
	return prd, nil
}
func (svc *service) List(page, limit int) ([]*domain.Product, error) {
	prd, err := svc.productRepo.List(page, limit)
	if err != nil {
		return nil, err
	}
	return prd, nil
}
func (svc *service) Update(p domain.Product) (*domain.Product, error) {
	prd, err := svc.productRepo.Update(p)
	if err != nil {
		return nil, err
	}
	return prd, nil
}
