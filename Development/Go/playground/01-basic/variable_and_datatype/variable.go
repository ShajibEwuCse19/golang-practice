package variable_and_datatype

import "fmt"

func VarAndDataType() {
	var a int = 10
	var b float32 = 10.5
	var c string = "Hello Shajib"
	var d bool = true
	var e complex64 = 10 + 5i

	fmt.Println("Integer value: ", a)
	fmt.Println("Float value: ", b)
	fmt.Println("String value: ", c)
	fmt.Println("Boolean value: ", d)
	fmt.Println("Complex value: ", e)
}
