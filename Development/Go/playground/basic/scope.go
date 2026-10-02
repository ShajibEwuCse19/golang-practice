package basic

import (
	"fmt"
)

func DisplayScope() { // function name starts with a capital letter, so it is exported and can be accessed from other packages.
	fmt.Println("Displaying scope...")

	x := 100
	y := 200

	fmt.Printf("Sum of local variables(%d, %d):%d\n", x, y, add(x, y))
	fmt.Printf("Sum of global variables(%d, %d):%d\n", a, b, add(a, b)) // Here, I can access the global variables a and b. So, the scope of global variables is throughout the package.
}

var (
	a = 10
	b = 20
)

func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func MyFunction() { // function name starts with a capital letter, so it is exported and can be accessed from other packages.
	//sum := add(x,y) // Here, I cannot access the local variables x and y. So, the scope of local variables is limited to the function in which they are declared.
	// x, y has no scope here. So, I cannot access them here.
	sum := add(a, b) // Here, I can access the global variables a and b. So, the scope of global variables is throughout the package.
	//a, b has scope here. So, I can access them here.
	fmt.Printf("MyFunction: Sum of global variables(%d, %d):%d\n", a, b, sum)
}
