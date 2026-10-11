package function

import "fmt"

func addTwoNum(a uint8, b uint8) (result uint8) { // result is the named return value
	result = a + b // no need to declare varialbe
	return // no need to return result, it'll return automatically
}

func DisplayNamedReturnValue() {
	fmt.Println("Displaying Named Return Value")
	fmt.Println("Result =", addTwoNum(10,20))
}