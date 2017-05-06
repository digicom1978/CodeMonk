package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	//"strings"
)


func countVowelsInString(str string) int {
	vowels := 0
	for v:=0; v<len(str); v++ {
		switch str[v] {
			case 'a',
			     'e',
			     'i',
			     'o',
			     'u',
			     'A',
			     'E',
			     'I',
			     'O',
			     'U':
			     //fmt.Println(str[v], 'a')
			     vowels++
		}
	}
	return vowels
}


func testCountVowelsInString() {
	fmt.Println("Enter no of lines:")

	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	numOfTestCases := s.Text()
	lines, _ := strconv.Atoi(numOfTestCases)
	//fmt.Printf("%d\n", lines)

	for i:=0; i<lines; i++ {
		s.Scan()
		str := s.Text()
		fmt.Printf("Num of vowels in %s is %d\n", str, countVowelsInString(str))
	}
}


func main() {
	fmt.Println("Begin …")

	testCountVowelsInString()
}