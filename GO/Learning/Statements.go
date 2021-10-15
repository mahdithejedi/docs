package main

import "fmt"

func main() {
	i := 3
	for ; i < 7; i++ {
		fmt.Println(i)
	}

	// there is no while but you can use `for` as while manner
	fmt.Printf("this")
	var j int = 0
	for {
		fmt.Printf(fmt.Sprint(j))
		if j == 3 {
			break
		}
		j++
	}

	// if else statement
	// it's like python if there is no parentheses
	l := 4
	if l > 3 {
		fmt.Printf("yes!")
	}
	// you can defind a variable in go
	// this variable is accessible from all branches
	if m := 4; m < 1 {
		fmt.Printf("NO")
	} else if m > 1 && m < 5 {
		fmt.Printf("YES")
	} else {
		fmt.Printf("NO")
	}
	// but you CAN'T access it out side of statement
	// for example this will cause en error fmt.Printf(fmt.Sprint(m))

	// switch case
	n := 2

	switch n {
	case 1, 2:
		fmt.Print("\n YES")
		fmt.Print("EXCATLY")
	case 3:
		fmt.Print("NO")
	case int:
		fmt.Print("YOU CAN declare types in cases!")
	default:
		fmt.Print("HI")
	}

}
