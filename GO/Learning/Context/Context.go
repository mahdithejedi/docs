package main

import (
	"fmt"
	"net/http"
	"time"
)

func Timeout(timeout time.Duration) <-chan time.Time {
	return time.After(timeout)
}

func Example(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Req received!")
	context := req.Context()
	deadlineTime, deadLineOk := context.Deadline()
	fmt.Println("deadline time is", deadlineTime, "deadline ok is", deadLineOk)
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "example")
	case <-context.Done():
		err := context.Err()
		fmt.Println("Server:", err)
	case <-Timeout(5 * time.Second):
		fmt.Fprintf(w, "timeout!!!")
	}
	fmt.Println("Done!")

}

func main() {
	http.HandleFunc("/example", Example)
	http.ListenAndServe(":5055", nil)
}
