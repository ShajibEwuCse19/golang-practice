package handlers

import (
	"ecom-project-improved/database"
	"ecom-project-improved/utils"
	"net/http"
	"slices"
	"strconv"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		utils.SendData(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for index,product := range database.ProductList {
		if product.ID == uint8(productID) {
			// database.ProductList = append(database.ProductList[:index], database.ProductList[index+1:]...)
			database.ProductList = slices.Delete(database.ProductList, index, index+1)
			utils.SendData(w, "Product deleted successfully", http.StatusOK)
			return
		}
	}

	utils.SendData(w, "Invalid product ID", http.StatusBadRequest)
}