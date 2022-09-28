package parser

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
)

type Link struct {
	Href string
	Text string
}

func Parser(r io.Reader) ([]Link, error){
	doc, err := html.Parse(r)
	if err != nil{
		return nil, err
	}
	fmt.Println(doc)
	return []Link{}, nil
}
