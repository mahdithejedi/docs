package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func _read_to_str(content io.ReadCloser) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(content)
	return buf.String()
}

func download_link(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	return _read_to_str(resp.Body), nil
}

func main() {
	fmt.Println(download_link("https://adventofcode.com/2021/day/1/input"))
}
