package services

import "ecom-project-improved/domain"

type service struct {
	productRepo ProductRepo
}

// Child(service) must need to implement all methods from parent(Service) if he want to return Service
// input-> productRepo , return Handler Service. [receive data from repo and return to parent Handler Service]
func NewService(productRepo ProductRepo) Service { // parent Service [port.go]
	return &service{ // child returning parent's parent [Service]
		productRepo: productRepo,
	}
}

// Child (service) implementing all the methods of parent(port.go/Service == handler/product/port.go) so that it can return Service [same signature]
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
