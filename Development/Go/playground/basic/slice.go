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

	// Slice Literal: An array without fixed size. 
	slice_literal := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice Literal =", slice_literal)
	fmt.Println("Address of slice_literal =", &slice_literal[0])
	fmt.Println("Capacity of slice_literal =", cap(slice_literal)) // capacity is 5
	fmt.Println("Pointer of slice_literal =", &slice_literal)
	fmt.Println("Length of slice_literal =", len(slice_literal)) // length is 5

	slice_literal = append(slice_literal, 100)
	fmt.Println("After appending the value Slice Literal =", slice_literal)
	fmt.Println("Address of slice_literal =", &slice_literal[0])
	fmt.Println("Capacity of slice_literal =", cap(slice_literal)) // capacity is 10, because in first time it takes 5 memory space, after that it takes 2 times more memory space
	fmt.Println("Pointer of slice_literal =", &slice_literal)
	fmt.Println("Length of slice_literal =", len(slice_literal)) // length is 6 [used memory space]

	slice_literal = append(slice_literal, 200, 300, 400, 500)
	fmt.Println("After appending the value Slice Literal =", slice_literal)
	fmt.Println("Address of slice_literal =", &slice_literal[0])
	fmt.Println("Capacity of slice_literal =", cap(slice_literal)) // capacity is 10 [used only blank memory space for storing values]
	fmt.Println("Pointer of slice_literal =", &slice_literal)
	fmt.Println("Length of slice_literal =", len(slice_literal)) // length is 10 [append 4 new values]

	slice_literal = append(slice_literal, 600, 700, 800, 900)
	fmt.Println("After appending the value Slice Literal =", slice_literal)
	fmt.Println("Address of slice_literal =", &slice_literal[0])
	fmt.Println("Capacity of slice_literal =", cap(slice_literal)) // capacity is 20, 2*previous memory = 2*10 = 20, this will work if the memory is full
	fmt.Println("Pointer of slice_literal =", &slice_literal)
	fmt.Println("Length of slice_literal =", len(slice_literal)) // length is 14 [append 4 new values]

	// Slice using make(): make([]type, length, capacity) but we can avoid the capacity
	// make([]type, length)
	slice_make := make([]int, 2, 5) // it will create a slice with 2 elements and capacity of 5, 
	fmt.Println("Slice using make() =", slice_make) // initially all values are zero
	fmt.Println("Address of slice_make =", &slice_make[0]) 
	fmt.Println("Capacity of slice_make =", cap(slice_make)) // capacity is 5
	fmt.Println("Pointer of slice_make =", &slice_make) 
	fmt.Println("Length of slice_make =", len(slice_make)) // length is 2 [used memory space]

	slice_make[0] = 100
	slice_make[1] = 200
	// slice_make[2] = 300 // this will give index out of range error. Because input using index depend on the length of the slice. 
	slice_make = append(slice_make, 300, 400, 500, 600, 700) // this will append the value if the memory is available
	fmt.Println("Slice using make() =", slice_make) 
	fmt.Println("Address of slice_make =", &slice_make[0]) 
	fmt.Println("Capacity of slice_make =", cap(slice_make)) // capacity is 10, because in first time it takes 5 memory space, after that it takes 2 times more memory space
	fmt.Println("Pointer of slice_make =", &slice_make) 
	fmt.Println("Length of slice_make =", len(slice_make)) // length is 7 [append 5 new values]

	// Empty slice or nil slice
	var nil_slice []int 
	fmt.Println("Nil slice =", nil_slice) // it will print [] and not [nil]
	// fmt.Println("Address of nil_slice =", &nil_slice[0]) // this will give index out of range error because the slice is nil, length is 0
	// fmt.Println("Capacity of nil_slice =", cap(nil_slice)) // this will give index out of range error
	// fmt.Println("Length of nil_slice =", len(nil_slice)) // this will give index out of range error

	//we can use only append()
	nil_slice = append(nil_slice, 100, 200, 300, 4, 5, 6, 7)
	fmt.Println("Nil slice =", nil_slice)
	fmt.Println("Address of nil_slice =", &nil_slice[0])
	fmt.Println("Capacity of nil_slice =", cap(nil_slice)) // 1 -> 2 -> 2 + 2 = 4 -> 4 + 2 = 6 -> 8 [append multiple elements increse memory +2 each time]
	fmt.Println("Length of nil_slice =", len(nil_slice)) // 7

	nil_slice = append(nil_slice, 200, 300)
	fmt.Println("After appending the value Nil slice =", nil_slice)
	fmt.Println("Address of nil_slice =", &nil_slice[0])
	fmt.Println("Capacity of nil_slice =", cap(nil_slice)) // 2 * 8 = 16, append single value increase memory 2 times. [8,16,32,64,128....]
	fmt.Println("Length of nil_slice =", len(nil_slice)) // 7 + 2 = 9

	// empty slice but not nil
	empty_slice := make([]int, 0)
	fmt.Println("Empty slice =", empty_slice) // it will print [] and not [nil]
	// fmt.Println("Address of empty_slice =", &empty_slice[0]) // this will give index out of range error because the slice is empty, length is 0
	fmt.Println("Capacity of empty_slice =", cap(empty_slice)) // 0
	fmt.Println("Length of empty_slice =", len(empty_slice)) // 0

}

/*
Slice:
1. Slice from an existing array.
2. Slice from an existing slice.
3. Slicing a slice.
4. Slice Literal.
5. Slice using make().
6. Empty slice or nil slice.

Slice has 3 property => pointer, length and capacity
Pointer => address of the first element of the slice [lower bound]
Length => number of elements in the slice [high-low]
Capacity => number of elements in the underlying array [len(arr)-low]
*/
