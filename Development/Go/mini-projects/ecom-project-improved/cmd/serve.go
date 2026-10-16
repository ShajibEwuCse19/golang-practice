package cmd

import (
	"ecom-project-improved/handlers"
	"ecom-project-improved/middleware"
	"ecom-project-improved/utils"
	"fmt"
	"net/http"
)

func Serve() {
	const port = ":8080"
	utils.FreePort(port) // Automatically check and kill any old or suspended process holding this port

	mux := http.NewServeMux() // router (multiplexer)

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))          // route (endpoint) and request handler with explicit HTTP method (GET).
	mux.Handle("POST /create-product", http.HandlerFunc(handlers.CreateProduct)) // route
	mux.Handle("GET /products/{id}", http.HandlerFunc(handlers.GetProductById)) // route
	mux.Handle("DELETE /products/{id}", http.HandlerFunc(handlers.DeleteProduct)) // route

	// here, mux.Handle expects a pattern and a http.Handler. So, I have created a middleware for handle CORS error and OPTION requests
	// The, GlobalRouterMiddleware checks all CORS related errors and finally return a http.Handler. So I passed it to mux.Handle. This is how middleware works in golang.
	middleware.GlobalRouterMiddleware(mux)

	fmt.Println("Server running on port: 8080")

	err := http.ListenAndServe(":8080", mux) // starts the server on port 3010 and listen for incoming requests.

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
