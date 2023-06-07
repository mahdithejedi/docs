package main

import "fmt"

//Find numbers devisable by 7 but not by 5

func SevenNotFive(_range int) []int {
	output := make([]int, 0)
	for i := 0; i < _range; i++ {
		if i%7 == 0 && i%5 != 0 {
			output = append(output, i)
		}
	}
	return output
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	output := SevenNotFive(40)
	printSlice(output)
}
