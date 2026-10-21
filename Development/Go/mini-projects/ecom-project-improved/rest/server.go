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
	//dependency
	productHandler *product.Handler
	userHandler    *user.Handler
	config         *config.Config
}

// inject the server's dependencies
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
	manager := middleware.NewManager() // create a new manager instance where initialized with an empty slice of global middlewares
	manager.AppendGlobalMiddlewares(
		middleware.Cors,
		middleware.Preflight,
		middleware.Logger,
	) // store all the common/global middlewares here

	mux := http.NewServeMux() // router (multiplexer)
	// here mux directly bind with server. So, when server get Preflight request (OPTIONS), it will directly call handler.
	// As, OPTION method has no handler, it will reject the request and browser will show Preflight error.
	// err := http.ListenAndServe(":8080", mux) // starts the server on port 3010 and listen for incoming requests.

	// Here is the solution of the problem: Cors(Preflight(Logger(mux)))
	wrapperMux := manager.WrapMuxWithGlobalMiddleware(mux) // added mux with mw before server connection

	//init routes
	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	err := http.ListenAndServe(":"+strconv.Itoa(server.config.HttpPort), wrapperMux) // bind the server port 8080 with wrapperMux[CorsWithPreflight(Logger(mux))]
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
