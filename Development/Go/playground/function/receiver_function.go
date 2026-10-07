package function

import "fmt"

type Person struct { // Person is a custom data type or struct
	Name string // member variable or field or attribute or property
	Age  int
}

// Greet is associated with the Person type. If we want to call the Greet function, we need to create an instance of the Person type 
// and then call the Greet function on that instance. This is known as a receiver function or method in Go.
func (p Person) Greet() { // Greet is a receiver function or method of type Person
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func DisplayReceiverFunction() {
	var person1 Person // Declare a variable of type Person
	person1 = Person{Name: "Alice", Age: 30} // person1 is a instance or object of type Person
	person1.Greet() // person1 is a instance of type Person. So, we can call the Greet function on person1. This is known as a receiver function or method in Go.


	person2 := Person{Name: "Bob", Age: 25}
	person2.Greet()

	var person3 Person 
	person3.Name = "Charlie"
	person3.Age = 35
	person3.Greet()
}