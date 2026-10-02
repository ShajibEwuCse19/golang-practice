package function

import "fmt"

//a := 10 //here,:=  is a short variable declaration, which Go only allows inside a function
var a int = 10 //here, we are declaring a variable 'a' of type int and initializing it with the value 10

func DisplayInitFunctionPackage() {
	fmt.Println("This is the standard function in the init function package.")
	fmt.Println("Value of a:", a) //25
}

// func init() { // computer calls `init` function automatically.
// 	fmt.Println("Init function in the function package is called.")
// 	fmt.Println("Value of a:", a) //10
// 	a += 5 // 10+5 = 15
// 	a += 10 // 15+10 = 25
// 	a := 100 // this a is a shadowing var of the global var a. This a is only accessible inside this init function.
// 	fmt.Println("Value of a inside init function:", a) //100
// 	a += 10
// 	fmt.Println("Value of a inside init function after adding 10:", a) //110
// }

// init function - This function is automatically called when the package is initialized. We can't call this function explicitly. It is used to set up package-level variables 
// or perform any necessary initialization tasks.