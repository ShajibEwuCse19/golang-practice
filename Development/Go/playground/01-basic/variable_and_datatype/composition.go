package variable_and_datatype

import (
	"fmt"
)

type Animal struct {
	Name string
}

func (a Animal) Speak() { // Method associated with the Animal struct
	fmt.Printf("%s makes a sound.\n", a.Name)
}

type Dog struct {
	//Animal Animal -> Composition using struct embedding
	Animal // Embedding the Animal struct directly
	Breed  string
}

func DisplayComposition() {

	// dog := Dog{
	// 	Animal: Animal{
	// 		Name: "Buddy",
	// 	},
	// 	Breed: "Golden Retriever",
	// }

	// dog := &Dog{
	// 	Animal: Animal{
	// 		Name: "Buddy",
	// 	},
	// 	Breed: "Golden Retriever",
	// }

	// dog := new(Dog)
	// dog.Name = "Buddy"
	// dog.Breed = "Golden Retriever"

	dog := &Dog{}
	dog.Name = "Buddy"
	dog.Breed = "Golden Retriever"

	dog.Speak() // Inherited method from Animal (Embedding methods)
	fmt.Printf("%s is a %s.\n", dog.Name, dog.Breed)

	ABC()
}

func ABC() {
	// a := A{Name: "A"}
	// b := B{A: a, Name: "B"}
	// c := C{B: b, Name: "C"}

	// fmt.Println("A Name:", a.Name)
	// fmt.Println("B Name:", b.Name)
	// fmt.Println("C Name:", c.Name)

	// // Accessing the embedded struct's field
	// fmt.Println("B's A Name:", b.A.Name)
	// fmt.Println("C's B Name:", c.B.Name)
	// fmt.Println("C's A Name:", c.B.A.Name)

	c := C{
		B: B{
			A: A{
				Name: "A",
			},
			Name: "B",
		},
		Name: "C",
	}

	fmt.Println("C Name:", c.Name)
	fmt.Println("C's B Name:", c.B.Name)
	fmt.Println("C's A Name:", c.B.A.Name)
}

type A struct {
	Name string
}

type B struct {
	A
	Name string
}

type C struct {
	B
	Name string
}

type Logger struct {} 
type Validator struct {}
type Repository struct {}

type Service struct {
    Logger
    Validator
    Repository
}

func ServiceMethod() {
	service := Service{
		Logger:    Logger{},
		Validator: Validator{},
		Repository: Repository{},
	}
	service.Log("This is a log message.")
	// service.Validate() // Assuming you have a Validate method in Validator
	// service.Save()     // Assuming you have a Save method in Repository
}

func (s Service) Log(message string) {
	fmt.Println("Logging:", message)
}
