package cmd

import (
	"ecom-project-improved/config"
	"ecom-project-improved/infra/db"
	"ecom-project-improved/repo"
	"ecom-project-improved/rest"
	srvProduct "ecom-project-improved/services/product"
	srvUser "ecom-project-improved/services/user"
	productHandler "ecom-project-improved/rest/handlers/product"
	userHandler "ecom-project-improved/rest/handlers/user"
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

	dbCon, err := db.NewConnection(config)
	if err != nil {
		fmt.Println("db connection err:", err)
		os.Exit(1)
	}

	// Migrate Database
	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println("db migration err:", err)
		os.Exit(1)
	}

	//repos
	userRepo := repo.NewUserRepo(dbCon)
	productRepo := repo.NewProductRepo(dbCon)

	//domains
	userSvc := srvUser.NewService(userRepo)
	productSvc := srvProduct.NewService(productRepo)

	userHandler := userHandler.NewHandler(userSvc)
	productHandler := productHandler.NewHandler(productSvc)

	server := rest.NewServer(productHandler, userHandler, config)
	server.StartServer()
}
