package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func input(input_string string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s \n", input_string)
	text, _ := reader.ReadString('\n')
	return text
}

func main() {
	out := input("hi!")
	out_slice := strings.Split(out, ",")
	for _, data := range out_slice {
		fmt.Printf("%s \n", data)
	}
}
