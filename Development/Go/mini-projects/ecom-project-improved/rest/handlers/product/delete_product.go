package product

import (
	"ecom-project-improved/database"
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

	database.DeleteProduct(uint8(productID))
	utils.SendData(w, "Product deleted successfully", http.StatusOK)
}