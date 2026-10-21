package cmd

import (
	"ecom-project-improved/config"
	"ecom-project-improved/infra/db"
	"ecom-project-improved/repo"
	"ecom-project-improved/rest"
	"ecom-project-improved/rest/handlers/product"
	"ecom-project-improved/rest/handlers/user"
	"ecom-project-improved/utils"
	"fmt"
	"os"
	"strconv"
)

func Serve() {
	config := config.GetConfig()
	fmt.Println("Version:", config.Version)
	fmt.Println("Service Name:", config.ServiceName)
	fmt.Println("Http Port:", config.HttpPort)
	utils.FreePort(":" + strconv.Itoa(config.HttpPort)) // Automatically check and kill any old or suspended process holding this port

	dbCon, err := db.NewConnection()
	if err != nil {
		fmt.Println("db connection err:", err)
		os.Exit(1)
	}
	userRepo := repo.NewUserRepo(dbCon)
	productRepo := repo.NewProductRepo()
	
	userHandler := user.NewHandler(userRepo)
	productHandler := product.NewHandler(productRepo)

	server := rest.NewServer(productHandler, userHandler, config)
	server.StartServer()
}
