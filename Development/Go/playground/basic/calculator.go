package basic

import "fmt"

func VarAndDataPractice() {
	fmt.Println("Variable and Data Type Practice")

	name := "Shajib"
	age := 25
	height := 5.7
	isEmployed := true

	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
	fmt.Println("Height: ", height)
	fmt.Println("Is Employed: ", isEmployed)

	age = 30
	fmt.Println("Updated Age: ", age)

	const country = "Bangladesh"
	fmt.Println("Country: ", country)

	fmt.Println("Sum of 50, 10 = ", Calculator(50, 10, "add"))
	fmt.Println("Subtraction of 50, 10 = ", Calculator(50, 10, "subtract"))
	fmt.Println("Multiplication of 50, 10 = ", Calculator(50, 10, "multiply"))
	fmt.Println("Division of 50, 10 = ", Calculator(50, 10, "division"))
	fmt.Println("Division of 0, 10 = ", Calculator(0, 10, "division"))
	fmt.Println("Division of 10, 0 = ", Calculator(10, 0, "division"))
	fmt.Println("Division of 0, 10 = ", Calculator(0, 10, "div"))

	intValue := 10
	floatValue := 5.5
	fmt.Println("result = ", intValue + int(floatValue))
	fmt.Println("result = ", float64(intValue) + floatValue)

	var number int
	var message string
	var isActive bool
	fmt.Println("Default value of int: ", number)
	fmt.Println("Default value of string: ", message)
	fmt.Println("Default value of bool: ", isActive)

	var sum int
	for i := 1; i <= 5; i++ {
		sum += i
	}
	fmt.Println("Sum of 1 to 5: ", sum)

}

func Calculator(a int, b int, operation string) int {
	switch operation {
	case "add":
		return a + b
	case "subtract":
		return a - b
	case "multiply":
		return a * b
	case "division":
		if b != 0 {
			return a / b
		} else {
			fmt.Println("Error: Divided by zero")
		}
	default:
		fmt.Println("Invalid operation")
	}
	return 0
}