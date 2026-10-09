package basic

import (
	"fmt"
)

func DisplayPointer() {
	age := 25 // age is a variable of type int

	p := &age // ampersand (&) => address of 
	fmt.Println("Address of age var:", p)
	value := *p // asterisk (*) => value at the address
	fmt.Println("Value of age var:", value) 

	//update value at the address
	age = 30 // directly update the value of the var age
	fmt.Println("Updated value of age var:", *p)
	*p = 35 // update the value from the address
	fmt.Println("Updated value of age var:", age)

	var ptr *int // ptr is a pointer variable of type int. It can hold the address of an int variable.
	ptr = &age

	fmt.Println("Age:", age)
	fmt.Println("Address:", ptr)
	fmt.Println("Value from the address:", *ptr)

	*ptr = 30
	fmt.Println("Updated Age:", age)

	persone := Person{
		Name: "Shajib",
		Age:  25,
	}
	fmt.Println("Person:", persone)
	fmt.Println("Person Name:", persone.Name)
	fmt.Println("Person Age:", persone.Age)

	personPtr := new(Person)
	personPtr.Name = "Rafi"
	personPtr.Age = 30
	fmt.Println("Person Pointer:", personPtr)
	fmt.Println("Person Pointer Name:", personPtr.Name)
	fmt.Println("Person Pointer Age:", personPtr.Age)


}

type Person struct {
	Name string
	Age  int
}

// Without pointer -> we need to pass billion size array values one by one, which is unrealistic time consuming. Pointer can solve the issue. 
// With pointer -> we will pass a billion size slice/array address [just a single number]. This is why we need to use pointer. 
func printArray(numbers *[]int) {
	fmt.Println(numbers) 
}
