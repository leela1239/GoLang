package main

import "fmt"

func main() {
	fmt.Println("Maps")

	var mp = map[int]string{} //initialization
	mp[1] = "Hi"
	fmt.Println(mp)

	mp1 := map[string]int{}
	mp1["a"] = 10
	mp1["b"] = 2
	fmt.Println(mp1)

	/*
		var m map[string]int // declaration
		m["a"] = 10 //panic: assignment to entry in nil map
		fmt.Println(m)
		fmt.Println("Len : ", len(m))  // Note : len() not gives a panic on nil map it gives 0
	*/ //Hence use make to craete a map

	mp2 := make(map[string]int)
	fmt.Println(mp2)
	fmt.Println("Len : ", len(mp2))

}
