package main

import (
	"HTMLLinkParser/parser"
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

func printUrl(urlChan chan string) {
	fmt.Println("started go routing")
	for url := range urlChan {
		fmt.Println("FUCKING urls is", url)
	}
	fmt.Println("OUT!")
}

type HtmlLinks struct {
	//url string
	fileData io.Reader
	wg       sync.WaitGroup
	urlChan  chan string
}

func main() {
	fileData, err := getFile()
	var wg sync.WaitGroup
	if err != nil {
		log.Fatal(err.Error())
	}
	urls := make(chan string)
	htmlLink := HtmlLinks{
		fileData: fileData,
		wg:       wg,
		urlChan:  urls,
	}
	err = parser.Parser(htmlLink)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer close(urls)
}

func getFile() (io.Reader, error) {
	fileName := flag.String("file", "", "file name to parse")
	flag.Parse()
	if strings.HasSuffix(*fileName, ".html") == false {
		errString := fmt.Sprintf("File %s should be html", *fileName)
		return nil, errors.New(errString)
	}
	if _, err := os.Stat(*fileName); errors.Is(err, os.ErrNotExist) {
		return nil, errors.New(fmt.Sprintf("File %s does not exists", *fileName))
	}
	fileData, err := os.Open(*fileName)
	if err != nil {
		return nil, err
	}
	//defer fileData.Close()
	return bufio.NewReader(fileData), nil

}
