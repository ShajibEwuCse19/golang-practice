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
	globalMiddlewares []Middleware // list of middlewares used with similar signature inside whole project will be stored here.
}

// NewManager is a constructor function that creates and returns a pointer to a new Manager struct
// Without return pointer => we need to return a new Manager struct object, so it will be copied. So it will be heavy and inefficient.
func NewManager() *Manager {
	return &Manager{ // returns a pointer of new instance of Manager struct
		globalMiddlewares: make([]Middleware, 0),
	}
}

// We can append middlewares like a chain with the use of this method.
// So that, we can use these middlewares in multiple places.
func (m *Manager) AppendMiddleware(middlewares ...Middleware) {
	m.globalMiddlewares = append(m.globalMiddlewares, middlewares...)
}

// Now, I'll create a receiver method for this Manager struct.
func (m *Manager) AddMiddleware(next http.Handler, middlewares ...Middleware) http.Handler { // ... is used to pass variadic arguments. which means we can pass 0 or list of arguments.
	handler := next
	// []middlewares = [Logger, GlobalRouterMiddleware, etc. whose signature is Middleware=func(http.Handler) http.Handler]
	// n = middleware.Logger(http.HandlerFunc(handlers.GetProducts))
	for _, mw := range middlewares {
		handler = mw(handler)
	}

	// wrap all the global middlewares => mw1(mw2(mw3(...(http.HandlerFunc(handlers.GetProducts)))))
	for _, mw := range m.globalMiddlewares {
		handler = mw(handler)
	}

	return handler // return the final handler
}

/*
 	handler = mw1(mw2(mw3(http.HandlerFunc(handlers.GetProducts))))
      = mw1(mw2(Logger(http.HandlerFunc(handlers.GetProducts))))
	  = mw1(AuthMiddleware(Logger(http.HandlerFunc(handlers.GetProducts))))
	  = GlobalRouterMiddleware(AuthMiddleware(Logger(http.HandlerFunc(handlers.GetProducts))))
*/
