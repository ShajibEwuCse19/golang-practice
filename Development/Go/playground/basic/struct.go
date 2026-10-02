package basic

import (
	"fmt"
)

type Address struct {
	City    string
	Country string
}

type Score struct {
	Subject string
	Score   map[string]int
}

type User struct {
	Name    string
	Age     int
	Email   string
	Address Address
	Skills  []string
	Score   Score
}

func UpdateEmail(user *User, newEmail string) {
	user.Email = newEmail
}

func (user User) DisplayStruct(message string) {
	fmt.Println("User: ", user)
	fmt.Println("Message: ", message)
	fmt.Println("User name =", user.Name)
	fmt.Println("User age =", user.Age)
	fmt.Println("User email =", user.Email)
	fmt.Println("User Address =", user.Address.City, user.Address.Country)

	UpdateEmail(&user, "xyz@gmail.com")
	fmt.Println("Updated User email:", user.Email)

	fmt.Println("User skills:", user.Skills)

	fmt.Println("User score subject:", user.Score.Subject)
	for name, score := range user.Score.Score {
		fmt.Printf("%s scored %d in %s.\n", name, score, user.Score.Subject)
	}

	averageScore := CalculateAverageScore(user.Score.Score)
	fmt.Printf("Average score in %s: %.2f\n", user.Score.Subject, averageScore)
}

func MainMethod() {
	user := User{
		Name:  "Shajib",
		Age:   25,
		Email: "abc@gmail.com",
		Address: Address{
			City:    "Dhaka",
			Country: "Bangladesh",
		},
		Skills: []string{
			"Go",
			"C++",
			"Java",
		},
		Score: Score{
			Subject: "Math",
			Score: map[string]int{
				"Shajib": 90,
				"Rafi":   85,
				"Rana":   95,
			},
		},
	}
	user.DisplayStruct("Hello, this is a test message!")
}

func CalculateAverageScore(scores map[string]int) float64 {
	if len(scores) == 0 {
		return 0.0
	}

	sum := 0.0
	for _, score := range scores {
		sum += float64(score)
	}
	return sum / float64(len(scores))
}
