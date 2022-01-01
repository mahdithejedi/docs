package main

import (
	"fmt"
	"math"
	"sort"
)

func SquaredMap(_range int) map[int]int {
	Squared := make(map[int]int)
	for i := 0; i < _range; i++ {
		Squared[i] = int(math.Pow(float64(i), 2))
	}
	return Squared
}

func printUnSortedMap(map_unsorted map[int]int) {
	fmt.Print("****** \nUNSORTED \n")
	for key, value := range map_unsorted {
		fmt.Printf("Key is %v, value is %v \n", key, value)
	}
	fmt.Print("******\n")
}

func PrintSortedMap(map_unsorted map[int]int) {
	keys := make([]int, len(map_unsorted))
	var i int = 0
	for key := range map_unsorted {
		keys[i] = key
		i++
	}
	sort.Ints(keys)
	fmt.Print("****** \nSORTED \n")
	for _, k := range keys {
		fmt.Printf("Key is %v, value is %v \n ", k, map_unsorted[k])
	}
	fmt.Print("******\n")
}

func main() {
	s := SquaredMap(12)
	printUnSortedMap(s)
	/*
		as we can see range in unsorted!
		so we have to convert them first to slice and then use them to point at the given map
		for more information: https://stackoverflow.com/a/23330883/9651641
	*/
	PrintSortedMap(s)

}
