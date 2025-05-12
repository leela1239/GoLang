package main

import "fmt"

func main() {

	//Array --> fixed size
	var arr [5]int // declaration not initialized , default values are assigned
	arr[1] = 1
	fmt.Println(arr)

	//declaration and initilaization
	var nums = [4]int{1, 2, 3, 4}
	nums[3] = 1
	fmt.Println(nums)

	//if we declare only dfew values remaining are set to set
	var num1 = [5]int{1, 2}
	num2 := num1
	num2[0] = 0
	fmt.Println("Num1 = ", num1)
	fmt.Println("Num2 = ", num2)

	//short notation
	arr1 := [2]int{}
	fmt.Println(arr1)

	//Note : THere is no concept of make

}
