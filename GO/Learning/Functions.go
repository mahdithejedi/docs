package main

import "fmt"

func multiple(a, b int) int {
	return a * b
}

// RETURN MULTIPLE VALUE

func multiple_2_return_value(a, b int) (int, int) {
	return multiple(a, 2), multiple(b, 2)
}

// VARIADIC FUNCTION
func sum(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// clouser functions or anonymous function
// you can defind a function in go without declaring it's name
func incSequance() func() int {
	initial_sequance := 0
	return func() int {
		initial_sequance++
		return initial_sequance
	}
}

func main() {
	res1, res2 := multiple_2_return_value(3, 6)
	fmt.Print(res1, res2)

	// sum of an array

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Print("\n \n sum is ", sum(numbers...))

	// clousers
	fmt.Print("\n") // just for add new line
	sequance := incSequance()
	fmt.Println("sould be 1", sequance())
	fmt.Println("sould be 1", sequance())
	fmt.Println("sould be 3", sequance())
	fmt.Println("sould be 4", sequance())
	// reset sequance
	sequance = incSequance()
	fmt.Println("override varibale")
	fmt.Println("sould be 1", sequance())
	fmt.Println("sould be 1", sequance())
	fmt.Println("sould be 3", sequance())
	fmt.Println("sould be 4", sequance())

	// declare a function
	var declare_func_var func(n int) int
	declare_func_var = func(n int) int {
		return n * n
	}
	fmt.Println("varibale is ", declare_func_var(3))

	new_declare := func(n int) int {
		return n * n
	}
	fmt.Println("NEW variable is ", new_declare(8))

}
