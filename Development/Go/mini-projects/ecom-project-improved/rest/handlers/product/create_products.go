package product

import (
	"ecom-project-improved/domain"
	"ecom-project-improved/utils"
	"encoding/json"
	"net/http"
)

type ReqCreateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	ImgURL      string  `json:"imgURL"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var reqProduct ReqCreateProduct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqProduct)
	if err != nil {
		http.Error(w, "Invalid request body. Give me valid json", 400)
		return
	}

	createdProduct, err := h.svc.Create(domain.Product{ // Map request DTO to domain model and call service contract
		Title: reqProduct.Title,
		Description: reqProduct.Description,
		Price: reqProduct.Price,
		ImgURL: reqProduct.ImgURL,
	})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	utils.SendData(w, createdProduct, http.StatusCreated)
}
