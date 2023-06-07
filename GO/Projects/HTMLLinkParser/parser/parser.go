package parser

import (
	"fmt"
	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func Parser(htmlLink stru) error {
	tokens := html.NewTokenizer(r)
	error := tokenize(tokens, urlsChain)
	if error != nil {
		return error
	}
	return nil
}

func getLink(tokens *html.Tokenizer, urlChain chan string) {
	key, value, moreAttrs := tokens.TagAttr()
	if string(key) == "href" {
		fmt.Println("Going to push " + string(value))
		urlChain <- string(value)
		fmt.Println("pushed!")
		return
	}
	if moreAttrs {
		getLink(tokens, urlChain)
	} else {
		return
	}
}

func tokenize(tokens *html.Tokenizer, urlChain chan string) error {
	for {
		tt := tokens.Next()
		switch tt {
		case html.ErrorToken:
			return tokens.Err()
		case html.StartTagToken, html.EndTagToken:
			tagName, hasAttrs := tokens.TagName()
			if string(tagName) == "a" && hasAttrs {
				getLink(tokens, urlChain)
			}
		default:
			continue
		}
	}
}
