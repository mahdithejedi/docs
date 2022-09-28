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
)


func main(){
	fileData, err := getFile()

	if err != nil {
		log.Fatal(err.Error())
	}
	err = parser.Parser(fileData)
	if err != nil{
		log.Fatal(err.Error())
	}
}

func getFile() (io.Reader, error){
	fileName := flag.String("file", "", "file name to parse")
	flag.Parse()
	if strings.HasSuffix(*fileName, ".html") == false{
		errString := fmt.Sprintf("File %s should be html", *fileName)
		return nil, errors.New(errString)
	}
	if _, err := os.Stat(*fileName); errors.Is(err, os.ErrNotExist){
		return nil, errors.New(fmt.Sprintf("File %s does not exists", *fileName))
	}
	fileData, err := os.Open(*fileName)
	if err != nil{
		return nil, err
	}
	//defer fileData.Close()
	return bufio.NewReader(fileData), nil


}
