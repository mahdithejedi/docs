package main

import "fmt"

type test struct {
	name string
}

var testSingleton *test

func New(name string) *test {
	if testSingleton == nil {
		testSingleton = &test{
			name: name,
		}
	}
	return testSingleton
}

func main() {
	first := New("Mahdi")
	Second := New("Reze")
	fmt.Println(first.name)
	fmt.Println(Second.name)
}
