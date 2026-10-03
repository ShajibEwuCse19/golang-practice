package function

import "fmt"

func DisplayClosure() {
	fmt.Println("This is the closure function in the function package.")
	call()
}

const a1 = 10
var a2 = 100

func outer() func() { // outer returns a function (closure) that has access to the variables in outer's scope
	money := 100 // outer-এর local variable
	age := 25

	fmt.Println("Age =", age)

	//Closure হলো এমন function যে নিজের বাইরের variable মনে রাখে এবং বদলাতে পারে, বাইরের function শেষ হয়ে গেলেও।
	// Closure = function + তার সাথে বহন করা একটা ছোট ব্যাগ, যার ভেতরে বাইরের variable-গুলো থাকে।
	show := func() { // show is the closure function that captures the variable money from outer's scope ||  এই outer function এর money ব্যবহার করছে
		money = money + a1 + a2 // a1, a2 are package-level variables, so they are accessible here and no need to capture them. but money is captured from outer's scope.
		// যেহেতু money var এই annonymous func এর বাইরের var, তাই Escape analysis এই money var কে heap এ রাখবে। তাই outer() শেষ হলেও money var বেঁচে থাকবে।
		//  আর a1, a2 program-এর পুরো সময় থাকে (global var), তাই তাদের জন্য heap allocation দরকার নেই।
		fmt.Println("Money =", money) 
	}
	return show // outer শেষ, কিন্তু money বেঁচে আছে
	// show নিজের ভেতরে money declare করেনি, বাইরের outer থেকে ধরেছে। এই "ধরে রাখা"-ই closure।
}

func call() {
	increment := outer()   // outer চলল: "Age" print, show stored in increment
	increment()            // show চলল but not outer()
	increment()            // show আবার চলল but not outer()

	increment2 := outer()  // outer আবার চলল: নতুন money, নতুন age and show stored in increment2
	increment2() 		  // show চলল but not outer()
	increment2() 		  // show আবার চলল but not outer()
}

/*
Output:
This is the closure function in the function package.
Age =  25
Money =  210
Money =  320
Age =  25
Money =  210
Money =  320
*/

/*
তিনটা গুরুত্বপূর্ণ শিক্ষা
১. Closure নিজের "ধরা" variable-এর state মনে রাখে
- increment প্রথমবার 210, দ্বিতীয়বার 320 দেয়, অর্থাৎ একই money variable প্রতিবার বাড়ছে। কপি নয়, সরাসরি reference।

২. প্রতিটা outer() call-এ আলাদা money
- increment2 শুরু করে আবার 210 থেকে, কারণ outer()-এর দ্বিতীয় call একটা নতুন money (100) বানিয়েছে। 
- increment আর increment2 দুটো আলাদা closure, প্রত্যেকের নিজস্ব money। একটা বদলালে অন্যটা প্রভাবিত হয় না।

৩. a1, a2 ধরা হয় না, money ধরা হয়
- money হলো outer-এর local variable, তাই show-কে সেটা capture করতে হয়।
- a1, a2 package-level, program-এর পুরো সময় থাকে, যেকোনো function সরাসরি পায়। Capture লাগে না।
- age closure-এ ব্যবহারই হয়নি, তাই capture হয়নি। outer শেষ হলে সেটা শেষ।
*/
