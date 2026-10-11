package function

import "fmt"

func addTwoNum(a uint8, b uint8) (result uint8) { // result is the named return value
	result = a + b // no need to declare varialbe
	return // no need to return result, it'll return automatically
}

func DisplayNamedReturnValue() {
	fmt.Println("Displaying Named Return Value")
	fmt.Println("Result =", addTwoNum(10, 20))

	//Defer with named return value
	a := calculate()
	fmt.Println("Main First =", a) // 0, 5, 15, 15(final result)
	
	// Defer with no named return value
	b := calc()
	fmt.Println("Main Second =", b) // 0, 5, 15, 5(final result)
}

func calculate() (result int) {
	fmt.Println("First =", result) // 0

	show := func() {
		result = result + 10 // 5 + 10 = 15
		fmt.Println("Defer =", result) // 15
	}
	defer show()

	result = 5
	fmt.Println("Second =", result) // 5

	return // return defer result 15
}

func calc() int {
	result := 0
	fmt.Println("First =", result) // 0

	show := func() { 
		result = result + 10 // 5 + 10 = 15
		fmt.Println("Defer =", result) // 15
	}
	defer show()

	result = 5
	fmt.Println("Second =", result) // 5

	return result // 5
}

/*
calc() and calculate() দুটোতেই defer, closure এবং result = result + 10 আছে। কিন্তু মূল পার্থক্য হলো calculate()-এ named return value ব্যবহার করা হয়েছে, 
আর calc()-এ সাধারণ return value ব্যবহার করা হয়েছে।

Output:
Displaying Named Return Value
First = 0
Second = 5
Defer = 15
Main First = 15

First = 0
Second = 5
Defer = 15
Main Second = 5

প্রশ্ন হলো: defer-এর মধ্যে দুটোতেই result-এর মান 15 হচ্ছে, তাহলে return value কেন আলাদা?

Answer:

calculate() function: এখানে result হলো named return value। func-এর signature-এই result নামের ভ্যারিয়েবলটি তৈরি হয়ে যায়।

defer inside calculate(): defer show( ) ফাংশনটি যখন চলে, তখন সে current result (অর্থাৎ 5) এর সাথে 10 যোগ করে 15 করে দেয়।

return in calculate(): যখন ফাংশনটি return হয়, তখন সে সেই মুহূর্তের result-এর ভ্যালু রিটার্ন করে। যেহেতু defer body-তে result পরিবর্তন হয়ে 15 হয়ে গেছে, তাই 15 রিটার্ন হবে।

calculate() — Named return    
result = 5    
    v
return — named return variable প্রস্তুত    
    v
defer চলে → result = 15    
    v
result return হয় 15

====================

calc() function: এখানে result নামের কোনো named return value নেই। এটি সাধারণ লোকাল ভ্যারিয়েবল।

defer inside calc(): defer show( ) এখানে result-এর মান 5 থেকে 15 করে দেয়।

return in calc(): কিন্তু return result স্টেটমেন্টটি result-এর মান 15 রিটার্ন করে না। কেন করে না?

Go-এর নিয়ম অনুযায়ী, when a named return value exists, return statement সেটির মান পরিবর্তন করার আগেই return হয়ে যায়। 
আর local variable (যেমন calc()-এর result) ফাংশন return হওয়ার সময় final value হিসেবে return হয়।

calc() — Unnamed return   
result = 5    
    v
return result → return value হিসেবে 5 evaluate হয়    
	v
defer চলে → local result = 15
    v
result return হয় 5


Finally,
পার্থক্য হচ্ছে return value-এর নিয়মে:
- calculate()-এ result নিজেই named return variable। defer সেটিকে পরিবর্তন করতে পারে।
- calc()-এ result সাধারণ local variable। return result-এর value আগে evaluate হয়ে যায়, তাই পরের পরিবর্তন return value-কে বদলায় না।

According to Go,
- Defer can update named return value
- But Defer can not update unnamed return value. Store the return value before defer block execute. So, return the stored value. 
*/