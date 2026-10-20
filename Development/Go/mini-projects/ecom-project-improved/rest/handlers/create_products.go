package handlers

import (
	"ecom-project-improved/database"
	"ecom-project-improved/utils"
	"encoding/json"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid request body. Give me valid json", 400)
		return
	}

	database.Store(newProduct)
	utils.SendData(w, newProduct, 201)
}
