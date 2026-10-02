package basic

import "fmt"

var globalVar int = 10 // This is a outer global variable

func DisplayVariableShadowing() {
	fmt.Println("Variable Shadowing Example")

	var localVar int = 20

	if localVar > 15 {
		var globalVar int = 30 // This globalVar shadows the outer globalVar
		// globalVar  = 30 // This modifies the outer globalVar
		fmt.Printf("Inner globalVar: %d\n", globalVar)
	}
	fmt.Printf("Outer globalVar: %d\n", globalVar)
}

// Here, we have two globalVar variables: one is defined outside the function (outer globalVar) and another is defined inside the if block (inner globalVar). 
// The inner globalVar shadows the outer globalVar within its scope.