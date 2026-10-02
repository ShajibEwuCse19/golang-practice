package basic

import "fmt"

func DisplayUserInput() {
	InputUserDate()
}

func InputUserDate() {
	var name string
	var age int

	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	fmt.Printf("Hello %s!\n You are %d years old.\n", name, age)

	var count int
	fmt.Print("Enter the number of items: ")
	fmt.Scan(&count)
	numberList := make([]int, count)

	for i := 0; i < count; i++ {
		fmt.Printf("Enter item %d: ", i+1)
		fmt.Scan(&numberList[i])
	}
	fmt.Println("\nNumbers: ", numberList)

	numberList = append(numberList, 100, 200, 300)
	fmt.Println("Updated Numbers: ", numberList)

	numbers := make([]int, 3, 5)
	numbers[0] += 100
	numbers = append(numbers, 200, 300, 400, 500)
	fmt.Println("Numbers: ", numbers)
}
