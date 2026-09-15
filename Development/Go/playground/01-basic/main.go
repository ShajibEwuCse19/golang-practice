package main

import (
	"fmt"
	"go-practice/variable_and_datatype"
)

func sum(a int, b int) int {
	return a + b
}

func display() {
	fmt.Println("Hello Shajib")
	fmt.Println("Hello World")
	fmt.Println("Hello Go")
	fmt.Println("Hello Playground")
}

func main() {
	// display()

	fmt.Println("Sum of two numbers = ", sum(5,5))

	// variable_and_datatype.VarAndDataType()
	variable_and_datatype.VarAndDataPractice()
}