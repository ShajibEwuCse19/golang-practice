package basic

import "fmt"

func DisplaySlice() {
	definitionOfASlice()
}

func definitionOfASlice() {
	arr := [6]string{"This", "is", "a", "go", "interview", "question"} // This is an array which is static and fixed size.

	fmt.Println("Address of arr =", &arr[0]) // this will give the address of the first element of the array

	// Slice is a part of an array
	// A slice has 3 property => pointer, length and capacity
	
	slice1 := arr[1:3]
	//here, pointer = 1, address of the starting value
	// length = end - start = 3 - 1 = 2
	// capacity = size of the main array - index of the starting pointer = len(arr) - pointer = 6 - 1 = 5

	fmt.Println("Slice-1 =", slice1)
	fmt.Println("Address of slice1 =", &slice1[0]) // this will give the address of the first element of the slice
	fmt.Println("Capacity of slice1 =", cap(slice1))// this will give the capacity of the slice
	fmt.Println("Pointer of slice1 =", &slice1)// this will give the address of the slice itself not the first element
	fmt.Println("Length of slice1 =", len(slice1))  // this will give the length of the slice

	// Slicing a slice
	slice2 := slice1[2:5] // pointer = 2, length = 5 - 2 = 3
	// here, slice2 slicing slice1 even if the length of slice1 is 2. This is possible because of the 
	// capacity of slice1 is 5.
	fmt.Println("Slice-2 =", slice2)
	fmt.Println("Address of slice2 =", &slice2[0])
	fmt.Println("Capacity of slice2 =", cap(slice2)) // capacity = len(slice1) - pointer of slice2 = 5 - 2 = 3
	fmt.Println("Pointer of slice2 =", &slice2)
	fmt.Println("Length of slice2 =", len(slice2))

	// arr index = 0 1 2 3 4 5
	// slice1    =   0 1 2 3 4
	// slice2    =       0 1 2 3 4, here, high value 3, 4 is not included in the slice1
	
}
