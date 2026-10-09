package basic

import "fmt"

func DisplaySlice() {
	definitionOfASlice()
}

func definitionOfASlice() {
	arr := [6]string{"This", "is","a","go","interview","question"} // This is an array which is static and fixed size.

	// Slice is a part of an array
	// A slice has 3 property => pointer, length and capacity
	slice1 := arr[1:3] 
	//here, pointer = 1, address of the starting value
	// length = end - start = 3 - 1 = 2
	// capacity = size of the main array - index of the starting pointer = len(arr) - pointer = 6 - 1 = 5
	fmt.Println("Slice-1 =", slice1)
}