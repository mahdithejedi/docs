package main

import (
	"context"
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

func withValue(w http.ResponseWriter, req *http.Request) {
	context1 := req.Context()
	context2 := context.WithValue(context1, "AUTH", "NON BARBARI")
	fmt.Println(context1.Value("AUTH"))
	fmt.Println(context2.Value("AUTH"))
	defer context2.Done()
	fmt.Println(context2.Value("AUTH"))
}

func withCancel(w http.ResponseWriter, req *http.Request) {
	context1, cancelFunc := context.WithCancel(req.Context())
	context1 = context.WithValue(context1, "KEY", "KEY IS KEY")
	fmt.Println("BEFORE", context1.Err())
	go func() {
		time.Sleep(15 * time.Second)
		// if you call this function context will be closed
		cancelFunc()
	}()
	select {
	case <-context1.Done():
		fmt.Println("DONE NONE IS")
	}
	fmt.Println("AFTER", context.Canceled)
	fmt.Println(context1.Value("KEY"), context1.Err())

}

func main() {
	http.HandleFunc("/example", Example)
	http.HandleFunc("/value", withValue)
	http.HandleFunc("/cancel", withCancel)
	http.ListenAndServe(":5055", nil)
}
