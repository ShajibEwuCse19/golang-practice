package handlers

import (
	"ecom-project-improved/database"
	"net/http"
	"ecom-project-improved/utils"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	data := database.List()
	if data == nil {
		utils.SendData(w, "No products found", 404)
		return
	}
	utils.SendData(w, data, 200)
}
