package main

import (
	"fmt"
)


func linerSearching(ar []int, target int) int {
	lastIdx := -1

	for i:=0; i<len(ar); i++ {
		if target == ar[i] {
			lastIdx = i
		}
	}
	return lastIdx
}

func testLinerSearching() {
	arr := make([]int, 5)
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 11
	arr[4] = 5

	fmt.Printf("Last Index of element[%d] is %d.\n", 1, linerSearching(arr,1))
}


func main() {
	fmt.Println("Begining…")
	testLinerSearching()
}