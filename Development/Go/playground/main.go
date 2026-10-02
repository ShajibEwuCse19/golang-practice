package main

import (
	"fmt"
	"go-practice/basic"
)

func main() {
	// Display()

	// fmt.Println("Sum of two numbers = ", sum(5,5))

	// basic.VarAndDataType()
	// basic.VarAndDataPractice()
	// basic.MainDisplayPractice2()
	// basic.DisplayUserInput()
	// basic.MainMethod()
	// basic.DisplayPointer()
	// basic.DisplayComposition()
	// basic.InterfaceExample()
	// basic.ErrorHandlingExample()
	basic.DisplayScope()
	basic.MyFunction()
}

func sum(a int, b int) int {
	return a + b
}

func Display() {
	// fmt.Println("Hello Shajib")
	// fmt.Println("Hello World")
	// fmt.Println("Hello Go")
	// fmt.Println("Hello Playground")

	// numbers := make([]int, 3, 5)
	// numbers[0] += 100
	// numbers = append(numbers, 200, 300, 400, 500)
	// fmt.Println("Numbers: ", numbers)

	// fmt.Printf("len(numbers) = %d\n", len(numbers))
	// fmt.Printf("cap(numbers) = %d\n", cap(numbers))

	// part := numbers[2:4]
	// fmt.Printf("%v\n", part)

	ages := map[string]int{
		"Shajib": 25,
		"Rafi":   30,
		"Rana":   35,
	}
	fmt.Println("Ages: ", ages)
	for name, age := range ages {
		fmt.Printf("%s is %d years old.\n", name, age)
	}

	ages["Shajib"] = 26
	ages["Rakib"] = 28
	fmt.Println("Updated Ages: ", ages)

	delete(ages, "Rafi")
	fmt.Println("Ages after deletion: ", ages)

	var heights map[string]float64 //nil map
	fmt.Println("Heights: ", heights)
	// heights["Shajib"] = 5.9
	// heights["Rafi"] = 6.0
	// heights["Rana"] = 5.8
	// fmt.Println("Updated Heights: ", heights)

	heights = make(map[string]float64)
	heights["Shajib"] = 5.9
	heights["Rafi"] = 6.0
	heights["Rana"] = 5.8
	fmt.Println("Updated Heights: ", heights)

}