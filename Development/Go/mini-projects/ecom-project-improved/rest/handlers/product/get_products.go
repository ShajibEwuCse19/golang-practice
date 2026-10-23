package product

import (
	"ecom-project-improved/utils"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	data, err := h.svc.List()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if data == nil {
		utils.SendData(w, "No products found", 404)
		return
	}
	utils.SendData(w, data, 200)
}
