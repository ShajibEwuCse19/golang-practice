package rest

import (
	"ecom-project-improved/config"
	"ecom-project-improved/rest/handlers/product"
	"ecom-project-improved/rest/handlers/user"
	"ecom-project-improved/rest/middleware"
	"fmt"
	"net/http"
	"strconv"
)

type Server struct {
	// Server dependencies
	productHandler *product.Handler
	userHandler    *user.Handler
	config         *config.Config
}

// NewServer injects dependencies into the Server instance.
func NewServer(
	productHandler *product.Handler,
	userHandler *user.Handler,
	config *config.Config,
) *Server {
	return &Server{
		productHandler: productHandler,
		userHandler:    userHandler,
		config:         config,
	}
}

func (server *Server) StartServer() {
	manager := middleware.NewManager() // Initialize middleware manager
	manager.AppendGlobalMiddlewares(
		middleware.Cors,
		middleware.Preflight,
		middleware.Logger,
	) // Register common global middlewares executed on every request

	mux := http.NewServeMux() // HTTP router (multiplexer)

	// Wrap mux router with global middlewares in proper execution order: Cors -> Preflight -> Logger -> Handler
	wrapperMux := manager.WrapMuxWithGlobalMiddleware(mux)

	// Initialize routes: registers handlers on the mux router with route-specific middlewares
	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	// Start listening for incoming HTTP requests on configured port
	err := http.ListenAndServe(":"+strconv.Itoa(server.config.HttpPort), wrapperMux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
