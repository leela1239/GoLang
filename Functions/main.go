package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func printEmpDet(id int, name string) (int, string) {
	return id, name
}

// Go's return values may be named
func findNatureOfNum(num int) (nature, div string) {
	if num >= 0 {
		nature = "Positive"
	} else {
		nature = "Negative"
	}
	if num%2 == 1 {
		div = "Odd"
	} else {
		div = "Even"
	}
	return //naked return function

}

// variadic function //Function with variable no.of arguments of same type
func Sum(nums ...int) int {
	fmt.Println("Len from variadic : ", len(nums))
	s := 0
	for _, val := range nums {
		s += val
	}
	return s
}

func main() {
	fmt.Printf("Sum value is : %v\n", add(2, 3))

	// A function can return multiple values
	val1, val2 := printEmpDet(101, "Leela")
	fmt.Println(val1, val2)

	//naked return function  // Go's return values may be named

	fmt.Println(findNatureOfNum(25))

	//variadic function
	fmt.Println("Sum = ", Sum(1, 2, 3))

	nums := []int{1, 2, 3, 4}
	fmt.Println("Sum = ", Sum(nums...)) // Sum(nums) don't work u need to pass individual elemanets not a slice , so nums... gives each element separately

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Sum = ", Sum(arr[:]...)) // convert array to slice first using [:] and follow thw above ...

}
