package cmd

import (
	"ecom-project-improved/handlers"
	"ecom-project-improved/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle(
		"GET /products",
		manager.AddMiddleware(http.HandlerFunc(handlers.GetProducts)),
	)

	mux.Handle(
		"POST /create-product",
		manager.AddMiddleware(http.HandlerFunc(handlers.CreateProduct)),
	)

	mux.Handle(
		"GET /products/{id}",
		manager.AddMiddleware(http.HandlerFunc(handlers.GetProductById)),
	)

	mux.Handle(
		"DELETE /products/{id}",
		manager.AddMiddleware(http.HandlerFunc(handlers.DeleteProduct)),
	)
}