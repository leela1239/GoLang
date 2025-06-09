package main

import "fmt"

type Stack struct {
	items []int
}

func (st *Stack) Push(val int) {
	st.items = append(st.items, val)
}

func (st *Stack) Pop() {
	if len(st.items) == 0 {
		fmt.Println("Stack is Empty")
	} else {
		val := st.items[len(st.items)-1]
		st.items = st.items[:len(st.items)-1]
		fmt.Printf("Popped value is : %d \n", val)
	}
}

func (st *Stack) Top() int {
	if len(st.items) == 0 {
		return -1
	} else {
		return st.items[len(st.items)-1]
	}
}

func main() {
	fmt.Println("Stack Using Struct")
	st := Stack{}
	st.Pop()
	fmt.Println(st.Top())
	st.Push(1)
	st.Push(2)
	st.Push(3)
	st.Pop()
	fmt.Println(st.Top())

}
