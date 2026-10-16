package main

import (
	"encoding/json" 
	"fmt"
	"net/http"
)

func createProduct(w http.ResponseWriter, r *http.Request) { 
	handleCors(w) 
	handlePreFlightRequest(w, r)

	var newProduct Product

	decoder := json.NewDecoder(r.Body) 
	err := decoder.Decode(&newProduct) 
	if err != nil {                   
		http.Error(w, "Invalid request body. Give me valid json", 400) 
		return
	}

	newProduct.ID = uint8(len(productList) + 1) 
	productList = append(productList, newProduct)

	sendData(w, newProduct, 201)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	handleCors(w) 	
	handlePreFlightRequest(w, r)
	sendData(w, productList, 200)
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int) { 
	w.WriteHeader(statusCode) 
	fmt.Fprintln(w, "Product created successfully")

	encoder := json.NewEncoder(w) 
	encoder.Encode(data)    
}

func handlePreFlightRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK) 
		return
	}
}

func handleCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*") 
	w.Header().Set("Content-Type", "application/json") 
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS") 
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")	
	w.Header().Set("Access-Control-Max-Age", "600") 
	w.Header().Set("Access-Control-Allow-Credentials", "true") 
}

func main() {
	mux := http.NewServeMux() // router (multiplexer)

	mux.Handle("GET /products", http.HandlerFunc(getProducts)) // route (endpoint) and request handler with explicit HTTP method (GET). 

	mux.Handle("POST /create-product", http.HandlerFunc(createProduct)) // route (endpoint) and request handler with explicit HTTP method (POST).

	fmt.Println("Server running on port: 8080")

	err := http.ListenAndServe(":8080", mux) // starts the server on port 3010 and listen for incoming requests. 

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}

type Product struct {
	ID          uint8   `json:"id"` 
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	ImgURL      string  `json:"imgURL"`
}

var productList []Product 

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Good condition sweet orange",
		Price:       50.60,
		ImgURL:      "https://www.oragneimage.com/imgOrange.png",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Good condition sweet Apple",
		Price:       150.20,
		ImgURL:      "https://www.appleimage.com/imgApple.png",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Grape",
		Description: "Good condition sweet Grape",
		Price:       90.60,
		ImgURL:      "https://www.grapeimage.com/imgGrape.png",
	}
	prd4 := Product{
		ID:          4,
		Title:       "Banana",
		Description: "Good condition sweet Banana",
		Price:       30.60,
		ImgURL:      "https://www.bananaimage.com/imgBanana.png",
	}
	prd5 := Product{
		ID:          5,
		Title:       "Mango",
		Description: "Good condition sweet Mango",
		Price:       120.60,
		ImgURL:      "https://www.mangoimage.com/imgMango.png",
	}

	productList = append(productList, prd1, prd2, prd3, prd4, prd5)
}
