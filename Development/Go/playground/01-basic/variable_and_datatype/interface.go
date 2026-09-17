package variable_and_datatype

import (
	"fmt"
)

type Speaker interface {
	talk() string
	Speak()
}

type SpeakingPerson struct {
	Name string
}

func (p SpeakingPerson) talk() string {
	return "Hello, my name is " + p.Name
}

func (p SpeakingPerson) Speak() {
	fmt.Println("I am a person.")
}

type SpeakingDog struct {
	Name string
}

func (d SpeakingDog) talk() string {
	return "Woof! My name is " + d.Name
}

func (d SpeakingDog) Speak() {
	fmt.Println("I am a dog.")
}

type SpeakingCat struct {} // it does not implement the Speaker interface

type SpeakingBird struct {
	Name string
}

func (b SpeakingBird) talk() string {
	return "Chirp! My name is " + b.Name
}

func (b SpeakingBird) Speak() {
	// empty implementation
}

func makeSpeak(s Speaker) {
	fmt.Println(s.talk())
	s.Speak()
}

func InterfaceExample() {
	person := SpeakingPerson{Name: "Shajib"}
	dog := SpeakingDog{Name: "Buddy"}
	// cat := SpeakingCat{}
	bird := SpeakingBird{Name: "Tweety"}

	makeSpeak(person)
	makeSpeak(dog)
	// makeSpeak(cat) // This will cause a compile-time error because SpeakingCat does not implement the Speaker interface
	makeSpeak(bird) // must implement all methods of the interface even if the method is empty
}