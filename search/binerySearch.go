package main

import (
	"fmt"
)


func binerySearch(ar []int, start int, end int, target int) int {
	mid := 0

	for ; start <= end;  {
		mid = (start + end) / 2
		if target < ar[mid] {
			end = mid-1
		} else if target > ar[mid] {
			start = mid+1
		} else {
			return mid
		}
	}

	return -1
}


func testBinerySearch() {
	//arr := make([]int, 10)
	//for i:=0; i<10; i++ {
	//	arr[i] = i+1
	//}
	arr := []int{1,2,4,5,7,8,11,12,17,20,21,24,26,27,28,31}


	target := 21
	targetIdx := binerySearch(arr, 0, len(arr)-1, target)

	fmt.Printf("Target %d is in [%d].\n", target, targetIdx)
}


func main() {
	fmt.Println("Begin …")
	testBinerySearch()
}