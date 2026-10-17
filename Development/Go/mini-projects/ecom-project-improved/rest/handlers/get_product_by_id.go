package handlers

import (
	"ecom-project-improved/database"
	"ecom-project-improved/utils"
	"net/http"
	"strconv"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id") //get the product ID from the URL parameters
	id, err := strconv.Atoi(productId) // convert the product ID from string to int

	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// search the product in the database using the ID
	for _, product := range database.ProductList { 
		if id == int(product.ID) {
			utils.SendData(w, product, http.StatusOK) // if found, send the product
			return
		}
	}

	utils.SendData(w, "Invalid product ID", http.StatusBadRequest) // if not found, send the error
}
