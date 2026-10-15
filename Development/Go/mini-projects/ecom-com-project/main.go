package main

import (
	"encoding/json" // Used to convert golang struct model to json format
	"fmt"
	"net/http"
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

	mux.HandleFunc("/create-product", createProduct) // route for create product

	fmt.Println("Server running on port: 3001")

	err := http.ListenAndServe(":3001", mux) //Starting a local server at port 3000. Listening for incoming requests if server port is free.

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
		http.Error(w, "Method not Allowed. Please send GET request.", 405) // Write the response. For error, with specific status code. Here, 400 which is a Bad Request.
		return
	}

	// Encode => Converting golang data structure to json format
	// creating json enconder and pass writer w
	// Task: Encode the productsList in json format and send it to the client (UI).
	encoder := json.NewEncoder(w) // create an encoder
	encoder.Encode(productList)// encode [struct to json format] productList and pass to w in json format for sending it to the client (UI).

}

func createProduct(w http.ResponseWriter, r *http.Request) { // client request to create a product - using POST and a request body. 
	w.Header().Set("Access-Control-Allow-Origin", "*") // call from any website, frontend - allow
	w.Header().Set("Content-Type", "application/json") // content will be json format

	if r.Method != http.MethodPost { // must be POST method to create a new product
		http.Error(w, "Method not Allowed. Please send POST request.", 405)
		return
	}

	// take the product details from Request body (r.Body) => title, description, price, imgURL. And, create an object of the Product type
	var newProduct Product

	//Decode => Converting json data to golang data structure
	// Decode the request body
	decoder := json.NewDecoder(r.Body) //creating a json decoder for the request body and decoder returns error if any
	err := decoder.Decode(&newProduct) // Decode [json to struct] of newProduct. Now newProduct is holding the data from the client.
	if err != nil { // any kind of error in found during decoding process [e.g. passing wrong type of data, or wrong format].
		http.Error(w, "Invalid request body. Give me valid json", 400) // Write the respons for error
		return
	}

	w.WriteHeader(201) // Write the response. For success, with specific status code. Here, 201 which is a Created.
	fmt.Fprintln(w, "Product created successfully")

	// Now I need to increase the ID of the new product (from the last product's ID) as client doesn't send the ID.
	newProduct.ID = uint8(len(productList) + 1) // current length + 1.

	// Append the new product to the productList.
	productList = append(productList, newProduct)
}

// product struct model for product
type Product struct {
	ID          uint8   `json:"id"` // Tag for converting ID to id in response body (key name in json)
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	ImgURL      string  `json:"imgURL"`
}

var productList []Product // products list

// created sample list for product using Product type
// init will call when the package is loaded and out products will be ready to use before create main()
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

	// store all the products in productList
	productList = append(productList, prd1, prd2, prd3, prd4, prd5)
}
