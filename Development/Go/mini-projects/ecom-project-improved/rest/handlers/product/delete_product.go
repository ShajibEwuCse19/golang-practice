package product

import (
	"ecom-project-improved/utils"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		utils.SendData(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	err = h.producRepo.Delete(uint8(productID))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	utils.SendData(w, "Product deleted successfully", http.StatusOK)
}
