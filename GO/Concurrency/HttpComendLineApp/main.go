package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Response struct {
	statusCode int
	url        string
}

func request(url string, Channel chan Response) {
	defer wg.Done()
	client := http.Client{
		Timeout: 2 * time.Second,
	}
	res, err := client.Get(url)
	var status Response
	if err != nil {
		status = Response{-1, url}
	} else {
		status = Response{
			statusCode: res.StatusCode,
			url:        url,
		}
	}
	Channel <- status
}

func main() {
	responses := make(chan Response)

	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <url1> <url2> ... <urln>")
	}
	for _, url := range os.Args[1:] {
		wg.Add(1)
		go request("https://"+url, responses)
	}
	for i := 0; i < len(os.Args)-1; i++ {
		respons := <-responses
		fmt.Printf("status code is %d for urls %s \n", respons.statusCode, respons.url)
	}

	wg.Wait()
	close(responses)
}
