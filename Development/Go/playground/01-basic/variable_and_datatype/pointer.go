package variable_and_datatype

import (
	"fmt"
)

func DisplayPointer() {
	age := 25
	var ptr *int 
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