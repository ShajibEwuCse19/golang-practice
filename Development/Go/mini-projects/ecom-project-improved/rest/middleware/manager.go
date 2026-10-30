package middleware

import (
	"net/http"
)

// Middleware represents a higher-order function that takes an http.Handler and returns a wrapped http.Handler.
type Middleware func(http.Handler) http.Handler

// Manager manages global and route-specific middleware chains.
type Manager struct {
	globalMiddlewares []Middleware
}

// NewManager initializes a new Manager instance with an empty slice of global middlewares.
func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

// AppendGlobalMiddlewares adds one or more middlewares to the global middleware chain.
func (m *Manager) AppendGlobalMiddlewares(middlewares ...Middleware) {
	m.globalMiddlewares = append(m.globalMiddlewares, middlewares...)
}

/*
মিডলওয়্যার চেইনিং নীতি:
স্লাইসে যেভাবে মিডলওয়্যার যুক্ত করা হবে (যেমন: CORS, Preflight, Logger),
রিকোয়েস্ট ঠিক সেই ক্রমানুসারে একটির পর একটি পার হয়ে মূল হ্যান্ডলারে পৌঁছাবে।
*/

// WrapMuxWithGlobalMiddleware wraps an http.Handler (e.g. ServeMux) with all registered global middlewares.
func (m *Manager) WrapMuxWithGlobalMiddleware(next http.Handler, middlewares ...Middleware) http.Handler {
	handler := next

	// Wrap in reverse order so the first middleware in the slice executes first:
	// e.g. Cors(Preflight(Logger(handler)))
	for i := len(m.globalMiddlewares) - 1; i >= 0; i-- {
		handler = m.globalMiddlewares[i](handler)
	}

	return handler
}

// AddMiddleware wraps a route handler with route-specific middlewares in execution order.
func (m *Manager) AddMiddleware(next http.Handler, middlewares ...Middleware) http.Handler {
	handler := next

	// Wrap in reverse order so middlewares[0] executes first
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return handler
}

/*
 	handler = mw1(mw2(mw3(http.HandlerFunc(handlers.GetProducts))))
      = mw1(mw2(Logger(http.HandlerFunc(handlers.GetProducts))))
	  = mw1(AuthMiddleware(Logger(http.HandlerFunc(handlers.GetProducts))))
	  = GlobalRouterMiddleware(AuthMiddleware(Logger(http.HandlerFunc(handlers.GetProducts))))
*/

/*
Summary of the code:
--------------------
type Middleware func(http.Handler) http.Handler -> create a Middleware type (function) which has same signature with http.Handler interface and all the middlewares in the project.
struct Manager -> create a struct for manager where list of []Middleware are stored
NewManager() -> initialize empty slice for globalMiddlewares
manager.AppendGlobalMiddlewares(middlewares...) -> append all the middlewares in the global list
manager.AddMiddleware(next, middlewares...) -> Pass next handler in all the Middlewares one by one and wrap them.
*/
