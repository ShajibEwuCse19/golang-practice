package user

import (
	"ecom-project-improved/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle(
		"POST /users",
		manager.AddMiddleware(
			http.HandlerFunc(h.CreateUser),
			middleware.DummyMiddleware,
		),
	)

	mux.Handle(
		"POST /users/login",
		manager.AddMiddleware(
			http.HandlerFunc(h.Login),
			middleware.DummyMiddleware,
		),
	)
}


