package main

import (
	"fmt"
)



func ternarySearch(ar []int, start int, end int, target int) int {
	if start <= end {
		mid1 := start + (end - start) / 3
		mid2 := end - (end - start) / 3

		if target == ar[mid1] {
			return mid1
		} else if target == ar[mid2] {
			return mid2
		} else if target < ar[mid1] {
			return ternarySearch(ar, start, mid1-1, target)
		} else if target > ar[mid2] {
			return ternarySearch(ar, mid2+1, end, target)
		} else if target > ar[mid2] {
			return ternarySearch(ar, mid1+1, mid2-1, target)
		}
	}

	return -1
}


func testTernarySearch() {
	arr := []int{2,3,5,6,8,9,12,13,14,23}

	target := 23
	idx := ternarySearch(arr, 0, len(arr)-1, target)
	fmt.Printf("Index of %d is %d\n", target, idx)
}


func main() {
	fmt.Println("Begin Ternary Search …")
	testTernarySearch()
}