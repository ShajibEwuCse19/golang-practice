package middleware

import (
	"net/http"
)

func GlobalRouterMiddleware(mux *http.ServeMux) http.Handler {
	handleAllReq := func(w http.ResponseWriter, r *http.Request) {
		handleCors(w)                // check permission for all header
		handlePreFlightRequest(w, r) // permission granted if the request method is OPTIONS (preflight request)

		mux.ServeHTTP(w, r) // now serve the request from the handler
	}

	corsHandler := http.HandlerFunc(handleAllReq)
	return corsHandler
}

func handleCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Max-Age", "600")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

func handlePreFlightRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
}
