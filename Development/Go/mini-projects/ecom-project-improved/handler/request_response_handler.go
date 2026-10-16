package handler

import (
	"fmt"
	"ecom-project-improved/productlist"
	"encoding/json"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	sendData(w, productlist.ProductList, 200)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct productlist.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid request body. Give me valid json", 400)
		return
	}

	newProduct.ID = uint8(len(productlist.ProductList) + 1)
	productlist.ProductList = append(productlist.ProductList, newProduct)

	sendData(w, newProduct, 201)
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int) { // data interface{} - it can take any type of data (int, float, struct, list etc).
	w.WriteHeader(statusCode)
	fmt.Fprintln(w, "Product created successfully")

	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}
