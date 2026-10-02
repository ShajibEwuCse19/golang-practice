package basic

import "fmt"

func CalculateTotal(price float64, quantity int) float64 {
	return price * float64(quantity)
} 

func CheckAge(age int) (string, bool) {
	if age >= 18 {
		return "You are eligible to vote.", true
	} 

	return "You are not eligible to vote.", false
}

func ScoresCalculator(scores []int) {
	fmt.Println("Scores:", scores)

	sum := 0
	for _, score := range scores {
		sum += score
	}
	fmt.Printf("Total score = %d\n", sum)
	fmt.Printf("Average score = %.2f\n", float64(sum)/float64(len(scores)))

	highest := scores[0]
	for i := 0; i < len(scores); i++ {
		if scores[i] > highest {
			highest = scores[i]
		}
	}
	fmt.Printf("Highest score = %d\n", highest)

	lowest := scores[0]
	for _, score := range scores {
		if score < lowest {
			lowest = score
		}
	}
	fmt.Printf("Lowest score = %d\n", lowest)

	for index, score := range scores {
		fmt.Printf("Score at index %d = %d\n", index, score)
	}
}

func MainDisplayPractice2() {
	fmt.Println("Variable and Data Type Practice")
	price := 10.5324
	quantity := 3
	total := CalculateTotal(price, quantity)
	fmt.Printf("Total price for %d items at $%.2f each is: $%.2f\n", quantity, price, total)

	age := 20
	category, isEligible := CheckAge(age)
	fmt.Printf("Age: %d - %s (Eligible: %t)\n", age, category, isEligible)

	ScoresCalculator([]int {10, 20, 30, 40, 50})

	scores := []int{85, 90, 78, 92, 88}
	ScoresCalculator(scores)
}