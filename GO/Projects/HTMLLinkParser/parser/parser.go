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

func Parser(r io.Reader) error {
	tokens := html.NewTokenizer(r)
	error := tokenize(tokens)
	if error != nil{
		return error
	}
	return nil
}

func tokenize(tokens *html.Tokenizer) error{
	for{
		tt := tokens.Next()
		switch tt{
		case html.ErrorToken:
			return tokens.Err()
		case html.TextToken:
			processTag(tokens.Text())
		case html.StartTagToken, html.EndTagToken:
			tn, _ := tokens.TagName()
			processTag(tn)
		default:
			continue
		}
	}
}

func processTag(tag []byte){
	fmt.Println(string(tag))
}