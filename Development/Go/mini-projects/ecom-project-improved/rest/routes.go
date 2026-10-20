package rest

import (
	"ecom-project-improved/rest/handlers"
	"ecom-project-improved/rest/middleware"
	"net/http"
)

func InitRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle(
		"GET /products",
		manager.AddMiddleware(
			http.HandlerFunc(handlers.GetProducts),
			middleware.DummyMiddleware,
		),
	)

	mux.Handle(
		"POST /create-product",
		manager.AddMiddleware(
			http.HandlerFunc(handlers.CreateProduct),
			middleware.DummyMiddleware,
			middleware.JwtAuthMiddleware,
		),
	)

	mux.Handle(
		"GET /products/{id}",
		manager.AddMiddleware(
			http.HandlerFunc(handlers.GetProductById),
			middleware.DummyMiddleware,
		),
	)

	mux.Handle(
		"DELETE /products/{id}",
		manager.AddMiddleware(
			http.HandlerFunc(handlers.DeleteProduct),
			middleware.DummyMiddleware,
		),
	)

	mux.Handle(
		"POST /users",
		manager.AddMiddleware(
			http.HandlerFunc(handlers.CreateUser),
			middleware.DummyMiddleware,
		),
	)

	mux.Handle(
		"POST /users/login",
		manager.AddMiddleware(
			http.HandlerFunc(handlers.Login),
			middleware.DummyMiddleware,
		),
	)
}
