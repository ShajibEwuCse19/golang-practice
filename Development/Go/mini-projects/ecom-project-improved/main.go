package main

import (
	"ecom-project-improved/cmd"
	"ecom-project-improved/config"
	"ecom-project-improved/utils"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Starting server...")

	config := config.GetConfig()
	fmt.Println("Version:", config.Version)
	fmt.Println("Service Name:", config.ServiceName)
	fmt.Println("Http Port:", config.HttpPort)

	utils.FreePort(":" + strconv.Itoa(config.HttpPort)) // Automatically check and kill any old or suspended process holding this port

	cmd.Serve()
}
