package handlers

import (
	"ecom-project-improved/database"
	"net/http"
	"ecom-project-improved/utils"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, database.ProductList, 200)
}
