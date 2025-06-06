package main

import (
	"fmt"
	"time"
)

func SayHello() {
	fmt.Println("Hello Leela")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("SayHello Completed")
}

func SayHi() {
	fmt.Println("Hi Leela")
	time.Sleep(2000 * time.Millisecond)
}

func main() {
	fmt.Println("Welcome!!")
	go SayHello()
	go SayHi()
	time.Sleep(1000 * time.Millisecond)
}
