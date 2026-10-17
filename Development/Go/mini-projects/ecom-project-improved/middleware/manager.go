package middleware

import (
	"net/http"
)

// type Middleware is a higher-order function that takes an http.Handler and returns an http.Handler
// Why type Middleware is needed?
// Multiple middlewares have similar signature => func(http.Handler) http.Handler.
// So, type Middleware type alias is used to avoid repetitive code.
type Middleware func(http.Handler) http.Handler

type Manager struct { // struct for manager
	globalMiddlewares []Middleware // list of Middleware (which is a function).
}

// Why NewManager Constructor is used?
//- Before using the Manager, we will call NewManager() for initialize empty slice for globalMiddlewares
func NewManager() *Manager { // initialize globalMiddlewares
	return &Manager{ // returns a pointer of new instance of Manager struct
		globalMiddlewares: make([]Middleware, 0),
	}
}

// Append all the mws in the mw global list
func (m *Manager) AppendGlobalMiddlewares(middlewares ...Middleware) { // 
	m.globalMiddlewares = append(m.globalMiddlewares, middlewares...)
}

/*
এটি নিশ্চিত করে যে স্লাইসে আপনি যেভাবে লিখবেন:
manager.AppendMiddleware(RecoverPanic, Logger, CORS)
রিকোয়েস্ট ঠিক সেভাবেই প্রথমে RecoverPanic $\to$ তারপর Logger $\to$ তারপর CORS পার হয়ে মূল হ্যান্ডলারে পৌঁছাবে।
*/

//Wrap all the middlewares in the handler.
func (m *Manager) AddMiddleware(next http.Handler, middlewares ...Middleware) http.Handler {
    handler := next

    // ১. Route Specific Middlewares রিভার্স অর্ডারে র‍্যাপ করা
    for i := len(middlewares) - 1; i >= 0; i-- {
        handler = middlewares[i](handler) // mw[i] means middleware (which is a function) at index i and mw[i](handler) means applying that middleware to the handler
    }

    // ২. Global Middlewares রিভার্স অর্ডারে র‍্যাপ করা
    for i := len(m.globalMiddlewares) - 1; i >= 0; i-- {
        handler = m.globalMiddlewares[i](handler)
    }

    return handler
}

/*
 	handler = mw1(mw2(mw3(http.HandlerFunc(handlers.GetProducts))))
      = mw1(mw2(Logger(http.HandlerFunc(handlers.GetProducts))))
	  = mw1(AuthMiddleware(Logger(http.HandlerFunc(handlers.GetProducts))))
	  = GlobalRouterMiddleware(AuthMiddleware(Logger(http.HandlerFunc(handlers.GetProducts))))
*/
