package function

import "fmt"

//standard function or named function
func add(a int, b int) int {
	return a + b
}

//anonymous function or IIFE (Immediately Invoked Function Expression)
func anonymousFunction() {
	func(a int, b int) {
		fmt.Println("Result from anonymous function:", add(a, b))
	}(2,3)
}

//add is a function expression
func functionExpression() {
	add := func(a int, b int) int {
		return a + b
	}

	fmt.Println("Result:", add(2, 3))
}

func DisplayFirstOrderFunction() {
	fmt.Println("This is the first order function in the function package.")
	addResult := add(2, 3)
	fmt.Println("Result from standard function:", addResult)
	functionExpression()
	anonymousFunction()
}

/*
First order function
- Standard function or named function.
- Anonymous function.
- IIFE (Immediately Invoked Function Expression).
- Function Expression. 
*/