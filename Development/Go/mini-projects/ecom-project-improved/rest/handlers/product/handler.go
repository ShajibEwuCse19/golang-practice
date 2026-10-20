package product

import "ecom-project-improved/repo"

type Handler struct {
	producRepo repo.ProductRepo
}

func NewHandler(producRepo repo.ProductRepo) *Handler {
	return &Handler{
		producRepo: producRepo,
	}
}