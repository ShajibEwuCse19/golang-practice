package cmd

import (
	"ecom-project-improved/config"
	"ecom-project-improved/middleware"
	"fmt"
	"net/http"
	"strconv"
)

func Serve() {
	manager := middleware.NewManager() // create a new manager instance where initialized with an empty slice of global middlewares
	manager.AppendGlobalMiddlewares(
		middleware.Cors,
		middleware.Preflight,
		middleware.Logger,
	) // store all the common/global middlewares here

	mux := http.NewServeMux() // router (multiplexer)

	fmt.Println("Server running on port: 8080")

	// here mux directly bind with server. So, when server get Preflight request (OPTIONS), it will directly call handler.
	// As, OPTION method has no handler, it will reject the request and browser will show Preflight error.
	// err := http.ListenAndServe(":8080", mux) // starts the server on port 3010 and listen for incoming requests.

	// Here is the solution of the problem: Cors(Preflight(Logger(mux)))
	wrapperMux := manager.WrapMuxWithGlobalMiddleware(mux) // added mux with mw before server connection
	initRoutes(mux, manager)

	err := http.ListenAndServe(":"+strconv.Itoa(config.GetConfig().HttpPort), wrapperMux)                                                         // bind the server port 8080 with wrapperMux[CorsWithPreflight(Logger(mux))]
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
