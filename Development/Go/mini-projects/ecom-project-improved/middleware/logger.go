package middleware

import (
	"fmt"
	"net/http"
	"time"
)

// Why next is http.Handler?
// because, Logger() is a middleware and it's will call before mux.handle() method.
// So, next of Logger() is mux.handle() method. Which is an http.Handler.
// After Logger() will finish its execution, it will call the next http.Handler()
// call Logger() first, next mux.handle() which expect http.Handler
func Logger(next http.Handler) http.Handler { //input next middleware's expected parameter
	return http.HandlerFunc( // after logging and execution [next handler -> request and response], it will return a http.Handler
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r) // call next [http.HandlerFunc(handlers.GetProducts)] http.Handler for execute the request
			end := time.Now()
			timeDiff := end.Sub(start)
			fmt.Printf("Logger: %s %s %s\n", r.Method, r.URL.Path, timeDiff.String())
		},
	)
}
