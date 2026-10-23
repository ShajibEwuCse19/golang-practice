package product

import (
	"ecom-project-improved/utils"
	"net/http"
	"strconv"
)

func (h *Handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id") //get the product ID from the URL parameters
	id, err := strconv.Atoi(productId) // convert the product ID from string to int

	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	data,err := h.svc.Get(uint8(id))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if data == nil {
		utils.SendData(w, "No products found", http.StatusBadRequest)
		return
	}
	utils.SendData(w, data, 200)
}
