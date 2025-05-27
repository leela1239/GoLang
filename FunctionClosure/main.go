package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	res := 0
	var a, b = 0, 1
	return func() int {
		res = a
		a, b = b, a+b
		return res
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	fmt.Println()
	f2 := fibonacci()
	fmt.Println(f2())
}
