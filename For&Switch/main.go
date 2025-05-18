package main

import "fmt"

func main() {

	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	// We don't have while loop the below syntx works as while loop
	sum := 1
	for sum < 25 {
		sum += sum
	}
	fmt.Printf("Sum is : %d \n", sum)

	//    Switch
}
