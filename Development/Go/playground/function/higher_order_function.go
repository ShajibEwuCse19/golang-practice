package function

import "fmt"

func sum(a int, b int) {
	fmt.Println("Sum:", a+b)
}

func square(a int) {
	fmt.Println("Square:", a*a)
}

// A higher-order function that takes another function as an argument
func completeOperation(a, b int, operation func(x int, y int)) { // here, operation is a function parameter. It's called callback function.
	operation(a,b)
}

func callSum() func (x int, y int) { // return type is a function that takes two integers as parameters
	return sum // returning the sum function
}

func DisplayHigherOrderFunction() {
	fmt.Println("Higher-order function example")

	completeOperation(200, 300, sum) // passing sum function as an argument to completeOperation
	// completeOperation(200, 300, square) // passing square function as an argument to completeOperation 
	// but it will not work because square function takes only one parameter where completeOperation expects a function that takes two parameters.


	/*
	Why sum() is passed as an argument to completeOperation()?
	- parameter of completeOperation() are two integers (a, b) and a function operation() with two integer parameters (x, y).
	- So, when I invoked/called completeOperation(), I need to pass two integers and a similar type function as arguments.
	- The sum() function is a standard function that takes two integers as parameters and returns their sum. 
	   It matches the signature of the operation parameter in completeOperation().
	- Therefore, I can pass sum() as an argument to completeOperation() to perform the addition operation on the provided integers (200 and 300).

	- completeOperation(200, 300, square) will not work because the signature of the square() function does not match the expected signature of the operation() 
	  parameter in completeOperation().
	*/

	// Calling a higher-order function that returns another function
	myOperation := callSum()
	myOperation(10, 20) // calling the returned function (sum) with arguments 10 and 20
}

// A higher-order function that takes another function as an argument and returns a function
func applyTwice(operation func(int) int) func(int) int { 
    return func(value int) int { // here, annonymous function is returned that takes an integer value as input. return type and the annonymous function's parameter type is similar.
        return operation(operation(value)) // applying the operation function twice to the input value
		/*
		এখানে আসল কাজ। লক্ষ্য করো, operation এই ভেতরের function-এর parameter নয় এবং এখানে declare-ও হয়নি। এটা এসেছে বাইরের applyTwice function থেকে। এটাই closure।
		কাজ করে ভেতর থেকে বাইরে:
		1. operation(value) প্রথমবার চলে -> operation()
		2. তার ফলাফল আবার operation-এ যায় -> operation(operation(value))
		3. দ্বিতীয় ফলাফল return হয় -> return result twice incremented value
		*/
    }
}

func DisplayFunctionInputAndOutput() {
    increment := func(value int) int { // anonymous function that increments a value by 1
        return value + 1
    }
	/*
	1. := হলো short variable declaration
	2. ডানদিকে একটা anonymous function, আর increment variable-এ সেই function-টাই রাখা হলো
	3. increment-এর type এখন func(int) int
	4. এখানে function চালানো হয়নি, শুধু variable-এ রাখা হয়েছে (এটাই "function is a value")
	*/

	fmt.Println("Increment once:", increment(5)) // Prints: 6 || increment(5) এবার function-টা চালায়: 5 + 1 = 6। Println সেটা print করে।

    incrementTwice := applyTwice(increment)
	/*
	1. increment function-টা operation parameter-এ পাঠানো হলো। এখানে increment(...) লেখা নেই, আছে শুধু increment, কারণ আমরা function-টাকে value হিসেবে দিচ্ছি, চালাচ্ছি না।
	2. applyTwice একটা নতুন function বানিয়ে ফেরত দেয়, আর সেটা incrementTwice-এ জমা হয়।
	3. তখনো কোনো গণনা হয়নি। এখন incrementTwice এমন একটা function, যে ভেতরে increment মনে রেখেছে।
	*/
	result := incrementTwice(5) // This will increment 5 twice, resulting in 7
	/*
	operation(operation(5))
	= increment(increment(5))
	= increment(6)
	= 7
	*/
    fmt.Println("Increment twice:", result) // Prints: 7
}

/*
Three conditions for higher-order function:
1. A function that takes one or more functions as arguments. || arguments or parameters -> another function
2. A function that returns another function as its result.
3. A function that does both of the above.
*/