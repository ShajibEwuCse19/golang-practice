package cmd

import (
	"ecom-project-improved/middleware"
	"ecom-project-improved/utils"
	"fmt"
	"net/http"
)

func Serve() {
	const port = ":8080"
	utils.FreePort(port) // Automatically check and kill any old or suspended process holding this port

	manager := middleware.NewManager()                                               // create a new manager instance where initialized with an empty slice of global middlewares
	manager.AppendGlobalMiddlewares(middleware.CorsWithPreflight, middleware.Logger) // store all the common/global middlewares here
	
	mux := http.NewServeMux()                                                        // router (multiplexer)

	initRoutes(mux, manager)

	fmt.Println("Server running on port: 8080")

	// here mux directly bind with server. So, when server get Preflight request (OPTIONS), it will directly call handler.
	// As, OPTION method has no handler, it will reject the request and browser will show Preflight error.
	// err := http.ListenAndServe(":8080", mux) 
	
	// Here is the solution of the problem:
	wrapperMux := manager.AddMiddleware(mux) // added mux with mw before server connection
	err := http.ListenAndServe(":8080", wrapperMux) // starts the server on port 8080 and listen for incoming requests.
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
