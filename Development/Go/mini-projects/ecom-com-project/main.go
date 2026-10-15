package main

import (
	"fmt"
	"net/http"
	"encoding/json" // Used to convert golang struct model to json format
)

// When router mux.HandleFunc is called -> It will create a route "/hello" and call helloHandler().
// Inside helloHandler() we have two arguments w -> Writer, r -> Request
// w will write the response ("Hello World") to the client (UI).
// r -> Request is used to read the request from the client (Not used here). e.g. login credentials username, password input from user from UI.
// For handling the request we pass this mux to the ListenAndServe function.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World") // this is the response.
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I am Shajib. I am a Software Engineer. Now I am learning Go")
}

func main() {
	mux := http.NewServeMux() // mux is a Router.

	mux.HandleFunc("/hello", helloHandler) // here pattern "/hello" is a Route of Router mux.

	mux.HandleFunc("/about", aboutHandler) // a route

	mux.HandleFunc("/products", getProducts) // route for get product list from getProducts()

	fmt.Println("Server running on port: 3006")

	err := http.ListenAndServe(":3006", mux) //Starting a local server at port 3000. Listening for incoming requests if server port is free.

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}


func getProducts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*") // allowing all origins to access the server. Skip CORS issues.
	// In production build, we should only allow the origin of the client.
	w.Header().Set("Content-Type", "application/json") //setting the content type to json

	// Request r has specifict http request Methods(GET/POST/PUT/DELETE,PATCH) and we check it's GET or not using r.Method
	//if r.Method != http.MethodGet // here, Method Get check it's "GET" or not

	if r.Method != "GET" { //checking GET method
		http.Error(w, "Please send GET request.", 400) // Write the response. For error, with specific status code. Here, 400 which is a Bad Request.
		return
	}

	// creating json enconder and pass writer w 
	// Task: Encode the productsList in json format and send it to the client (UI).
	encoder := json.NewEncoder(w) // encoder
	encoder.Encode(productList) // encode productList and pass to response writer w in json format for sending it to the client (UI).
	
}

// product struct model for product
type Product struct {
	ID          uint8  `json:"id"` // Tag for converting ID to id in response body (key name in json)
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       float32 `json:"price"`
	ImgURL      string `json:"imgURL"`
}

var productList []Product // products list

// created sample list for product using Product type
// init will call when the package is loaded and out products will be ready to use before create main()
func init() {
	prd1 := Product{
		ID: 1,
		Title: "Orange",
		Description: "Good condition sweet orange",
		Price: 50.60,
		ImgURL: "https://www.oragneimage.com/imgOrange.png",
	}
	prd2 := Product{
		ID: 2,
		Title: "Apple",
		Description: "Good condition sweet Apple",
		Price: 150.20,
		ImgURL: "https://www.appleimage.com/imgApple.png",
	}
	prd3 := Product{
		ID: 3,
		Title: "Grape",
		Description: "Good condition sweet Grape",
		Price: 90.60,
		ImgURL: "https://www.grapeimage.com/imgGrape.png",
	}
	prd4 := Product{
		ID: 4,
		Title: "Banana",
		Description: "Good condition sweet Banana",
		Price: 30.60,
		ImgURL: "https://www.bananaimage.com/imgBanana.png",
	}
	prd5 := Product{
		ID: 5,
		Title: "Mango",
		Description: "Good condition sweet Mango",
		Price: 120.60,
		ImgURL: "https://www.mangoimage.com/imgMango.png",
	}

	// store all the products in productList
	productList = append(productList, prd1, prd2, prd3, prd4, prd5)
}
