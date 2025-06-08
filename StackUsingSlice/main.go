package main

import "fmt"

// Stack implemented using slice
type Stack []int

// push
func (s *Stack) Push(val int) {
	*s = append(*s, val)
}

// pop
func (s *Stack) Pop() (int, string) {
	if len(*s) == 0 {
		return -1, "Stack is Empty"
	}
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val, ""
}

// top
func (s *Stack) Top() int {
	val := (*s)[len(*s)-1]
	return val
}

func (s *Stack) PrintStack() {

	for _, val := range *s {
		fmt.Print(val, " ")
	}
	fmt.Println()
}

func main() {
	fmt.Println("Stack")

	st := Stack{}
	val, err := st.Pop()
	fmt.Print("Pop : ")
	if err != "" {
		fmt.Println("Stack is Empty")
	} else {
		fmt.Println(val)
	}
	st.Push(1)
	st.Push(2)
	st.Push(3)
	st.PrintStack()
	fmt.Println("Top : ", st.Top())
	val1, err1 := st.Pop()
	fmt.Print("Pop : ")
	if err1 != "" {
		fmt.Println("Stack is Empty")
	} else {
		fmt.Println(val1)
	}
	fmt.Println("Top : ", st.Top())
	st.PrintStack()

}

