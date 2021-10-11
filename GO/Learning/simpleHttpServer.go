package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request){
	var name, _ = os.Hostname()
	fmt.Fprintf(w, "<h1> Request have been sended from %s <h1/>", name)
}

func main(){
	var port int = 80
	fmt.Fprintf(os.Stdout, "webserver started on port %s", port)
	http.Httpfunc('/', handler)
	http.ListenAndServer(":" + port, nil)
}