package function

import "fmt"

func AnonymousFunctionDisplay() {
	fmt.Println("This is an anonymous function in the function package.")

	//anonymous function => Immediately Invoked Function Expression (IIFE)
	func(a,  b int) { // no function name. Inside the function body is called function definition
		sum := a + b
		fmt.Println("Sum of two numbers from anonymous function = ", sum)
	}(5, 5) // immediately calling/invoking/executing the anonymous function.  call = invoke = execute, all are same meaning

	// Function expression => Assigning a function to a variable
	sumFunc := func(a, b int) int { // function expression. Assigning a function to a variable
		return a + b
	}
	fmt.Println("Sum from function expression = ", sumFunc(5, 5))
}

// anonymous function is a function with out a name. It can be defined and called inline, often used for short-lived operations or as arguments to higher-order functions.  