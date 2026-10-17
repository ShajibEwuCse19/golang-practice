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

	mux := http.NewServeMux()                   // router (multiplexer)
	manager := middleware.NewManager()          // create a new manager instance where initialized with an empty slice of global middlewares
	manager.AppendGlobalMiddlewares(middleware.CorsWithPreflight, middleware.Logger) // store all the common/global middlewares here

	initRoutes(mux, manager)

	fmt.Println("Server running on port: 8080")

	err := http.ListenAndServe(":8080", mux) // starts the server on port 3010 and listen for incoming requests.

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
