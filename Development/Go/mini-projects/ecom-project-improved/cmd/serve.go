package cmd

import (
	"ecom-project-improved/config"
	"ecom-project-improved/rest"
	"ecom-project-improved/rest/handlers/product"
	"ecom-project-improved/rest/handlers/user"
	"ecom-project-improved/utils"
	"fmt"
	"strconv"
)

func Serve() {
	config := config.GetConfig()
	fmt.Println("Version:", config.Version)
	fmt.Println("Service Name:", config.ServiceName)
	fmt.Println("Http Port:", config.HttpPort)

	utils.FreePort(":" + strconv.Itoa(config.HttpPort)) // Automatically check and kill any old or suspended process holding this port

	productHandler := product.NewHandler()
	userHandler := user.NewHandler()

	server := rest.NewServer(productHandler, userHandler, config)
	server.StartServer()
}
