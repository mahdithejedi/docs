package parser

import "fmt"

func ProcessTag(tag []byte) {
	fmt.Println(string(tag))
}

func ProcessLink(link string) {
	fmt.Println("Link is " + link)
}
