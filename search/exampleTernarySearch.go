package main

import (
	"fmt"
)


func fnDouble(x float64) float64 {
	return -1*x*x + 2*x + 3
}


func findMaxByTernarySearch(start float64, end float64) (float64, float64) {

	l1 := start
	l2 := end
	var x float64 = 0

	for i:=0; i<200; i++ {
		mid1 := (2*l1 + l2) / 3
		mid2 := (2*l2 + l1) / 3

		if fnDouble(l1) > fnDouble(l2) {
			l2 = mid2
		} else {
			l1 = mid1
		}
	}

	x = l1
	return fnDouble(x), x

}


func testExampleTernarySearch() {
	max, idx := findMaxByTernarySearch(-2, 6)
	fmt.Printf("max: %f, index: %f\n", max, idx)

	max, idx = findMaxByTernarySearch(-104, 0)
	fmt.Printf("max: %f, index: %f\n", max, idx)
}


func main() {
	fmt.Println("Begin example of ternary search …")

	testExampleTernarySearch()
}