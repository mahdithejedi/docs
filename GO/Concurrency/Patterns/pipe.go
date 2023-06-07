package main

import "fmt"

// __source__ = 'https://go.dev/blog/pipelines'

func number(nums ...int) <-chan int {
	num := make(chan int)
	sender := func() {
		for _, i := range nums {
			num <- i
		}
		close(num)
	}
	go sender()
	return num
}

func square(nums <-chan int) <-chan int {
	sq := make(chan int)
	go func() {
		for i := range nums {
			sq <- i * i
		}
		close(sq)
	}()
	return sq
}

func main() {
	// Pipe pattern
	for value := range square(number(12, 34, 56)) {
		fmt.Println(value)
	}

}
