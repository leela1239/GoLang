package main

import "fmt"

var klh = 9113 // can declare variables otside of function

//klm := 12 // short assignment statement not works outside the function

func main() {

	//single value with data type mentioned
	var num int = 9
	fmt.Println(num)

	//single valuewith data type mentioned
	var x = 9
	fmt.Println(x)

	//multiple values with same datatype
	var num1, num2 int = 1, 2
	fmt.Println(num1, num2)

	//multiple values if datatype not specified no need to be of same datatype
	var a, b, c = 1, 2, "Hi"
	fmt.Println(a, b, c)

	//multiple values datatype mentioned with multiple datatypes
	var a1 int = 1
	var a2 string = "hii"
	fmt.Println(a1, a2)

	//multiple values datatype not mentioned with multiple datatypes
	var (
		p = 1
		q = "Hello"
	)
	fmt.Println(p, q)

	//short assignment statement works only inside the function
	newVal := 35
	fmt.Println("new val is : ", newVal)

	newVal = 20
	fmt.Println("new val is : ", newVal)

	fmt.Println("klh value is : ", klh)

}
