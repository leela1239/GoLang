package main

import "fmt"

func appendFun(s []int) {
	s = append(s, 100, 200)
	s[0] = 9
	fmt.Println("s = ", s)
}

func main() {

	// var arr = [5]int{1, 2, 3, 4, 5}
	// fmt.Println(arr)

	// /*
	// 	var nums1 = []int{}
	// 	nums1[0] = 1
	// 	fmt.Println(nums1)
	// */ //This code gives panic  as nums1 is pointed to nil as the values are not initiliasised

	// //so it is better to use make
	// nums := make([]int, 3)
	// nums[0] = 1
	// fmt.Println("nums = ", nums)

	// var slice1 = arr[:3]
	// //slice1 = append(slice1, 25)
	// fmt.Println("slice1 : ", slice1)
	// fmt.Println("Len :", len(slice1))
	// fmt.Println("Cap : ", cap(slice1))
	// fmt.Println("arr : ", arr)

	// var slice2 = arr[:5]
	// slice2[1] = 11
	// fmt.Println(slice2)

	// //modifying slice changes the array
	// fmt.Println(arr)

	s := make([]int, 0, 5)
	fmt.Println("s1 : ", s)
	s = append(s, 1, 2)
	fmt.Println("s2 : ", s)
	appendFun(s)
	s = append(s, 11, 22)
	appendFun(s)
	fmt.Println("s3 :", s)

}
