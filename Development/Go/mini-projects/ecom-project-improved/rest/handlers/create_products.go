package handlers

import (
	"ecom-project-improved/database"
	"encoding/json"
	"net/http"
	"ecom-project-improved/utils"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid request body. Give me valid json", 400)
		return
	}

	newProduct.ID = uint8(len(database.ProductList) + 1)
	database.ProductList = append(database.ProductList, newProduct)

	utils.SendData(w, newProduct, 201)
}
