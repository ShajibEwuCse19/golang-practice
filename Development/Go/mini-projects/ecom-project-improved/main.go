package main

import (
	"fmt"
	"net/http"
	"ecom-project-improved/middleware"
	"ecom-project-improved/freeport"
	"ecom-project-improved/handler"
)

func main() {
	const port = ":8080"
	freeport.FreePort(port) // Automatically check and kill any old or suspended process holding this port

	mux := http.NewServeMux() // router (multiplexer)

	// here, mux.Handle expects a pattern and a http.Handler. So, I have created a middleware for handle CORS error and OPTION requests
	// The, HandleCorsMiddleware checks all CORS related errors and finally return a http.Handler. So I passed it to mux.Handle. This is how middleware works in golang.
	mux.Handle("GET /products", middleware.HandleCorsMiddleware(http.HandlerFunc(handler.GetProducts))) // route (endpoint) and request handler with explicit HTTP method (GET). 
	mux.Handle("POST /create-product", middleware.HandleCorsMiddleware(http.HandlerFunc(handler.CreateProduct))) // route (endpoint) and request handler with explicit HTTP method (POST).

	fmt.Println("Server running on port: 8080")

	err := http.ListenAndServe(":8080", mux) // starts the server on port 3010 and listen for incoming requests. 

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}

