package main

import (
	"fmt"
)

var top int = -1

func push(stack []int, x int, n int) {
	if top == n-1 {
		fmt.Println("Size is full, overflow")
	} else {
		top++
		stack[top] = x
	}
}

func pop(stack []int) int {
	if top < 0 {
		fmt.Println("This stack is empty …")
		return -1
	} else {
		temp := top
		top-=1
		return stack[temp]
	}
}

func printStack(ar []int) {
	end := top
	for i:=end; i>=0; i-- {
		fmt.Println(ar[i])
	}
	fmt.Println("The end of stack ...")
}


func isEmpty(ar []int) bool {
	if top == -1 {
		return true
	} else {
		return false
	}
}


func testStack() {
	arr := make([]int, 10)
	printStack(arr)
	push(arr, 10, len(arr))
	push(arr, 11, len(arr))
	push(arr, 12, len(arr))
	printStack(arr)
	fmt.Printf("Poped: %d\n", pop(arr))
	printStack(arr)
	fmt.Printf("Poped: %d\n", pop(arr))
	fmt.Printf("Poped: %d\n", pop(arr))
	fmt.Printf("Poped: %d\n", pop(arr))

	fmt.Printf("Is empty? %t\n", isEmpty(arr))
}


func main() {
	fmt.Println("Begin ...")
	testStack()
}