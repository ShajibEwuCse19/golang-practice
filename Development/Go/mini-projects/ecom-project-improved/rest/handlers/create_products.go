package handlers

import (
	"ecom-project-improved/database"
	"ecom-project-improved/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	if !checkAccessToken(w, r) {
		return
	}
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

func checkAccessToken(w http.ResponseWriter, r *http.Request) bool {
	header := r.Header.Get("Authorization")
	if header == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return false
	}

	headerArray := strings.Split(header, " ")
	if len(headerArray) != 2 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return false
	}

	accessToken := headerArray[1] //header.payload.signature

	if utils.VerifyJwt(accessToken) {
		fmt.Println("Token is valid")
		fmt.Println("Token: ", accessToken)
		return true
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return false
	}
} 
