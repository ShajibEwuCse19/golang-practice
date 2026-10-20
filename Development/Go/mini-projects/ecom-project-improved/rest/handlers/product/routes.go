package product

import (
	"ecom-project-improved/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle(
		"GET /products",
		manager.AddMiddleware(
			http.HandlerFunc(h.GetProducts),
			middleware.DummyMiddleware,
		),
	)

	mux.Handle(
		"POST /create-product",
		manager.AddMiddleware(
			http.HandlerFunc(h.CreateProduct),
			middleware.DummyMiddleware,
			middleware.JwtAuthMiddleware,
		),
	)

	mux.Handle(
		"GET /products/{id}",
		manager.AddMiddleware(
			http.HandlerFunc(h.GetProductById),
			middleware.DummyMiddleware,
		),
	)

	mux.Handle(
		"DELETE /products/{id}",
		manager.AddMiddleware(
			http.HandlerFunc(h.DeleteProduct),
			middleware.DummyMiddleware,
		),
	)
}