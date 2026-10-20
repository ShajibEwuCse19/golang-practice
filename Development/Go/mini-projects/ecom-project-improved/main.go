package main

import (
	"ecom-project-improved/cmd"
	"fmt"
)

func main() {
	fmt.Println("Starting server...")
	cmd.Serve()
}

// using JWT from utils

// import (
// 	"ecom-project-improved/utils"
// 	"fmt"
// )

// func main() {
// 	myJwt, err := utils.CreateJwt(
// 		"my-secret",
// 		utils.PayLoad{
// 			Sub:         "101",
// 			FirstName:   "Md Shajibul",
// 			LastName:    "Islam",
// 			Email:       "shajib@gmail.com",
// 			IsShopOwner: false,
// 		},
// 	)
// 	if err != nil {
// 		fmt.Println("Error: ", err.Error())
// 	} else {
// 		fmt.Println("JWT: ", myJwt)
// 	}
// }
