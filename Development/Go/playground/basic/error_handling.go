package basic

import (
	"fmt"
	"errors"
)

func ErrorHandlingExample() {
	result, err := divide(10, 0)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)
}

type error interface {
	Error() string
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}