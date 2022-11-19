package main

// __source__ = 'https://go.dev/blog/context'

import (
	"Google/google"
	"Google/userip"
	"context"
	"html/template"
	"log"
	"net/http"
	"time"
)

var resultsTemplate = template.Must(template.New("results").Parse(`
<html>
<head/>
<body>
  <ol>
  {{range .Results}}
    <li>{{.Title}} - <a href="{{.URL}}">{{.URL}}</a></li>
  {{end}}
  </ol>
  <p>{{len .Results}} results in {{.Elapsed}}; timeout {{.Timeout}}</p>
</body>
</html>
`))

type TemplateStruct struct {
	Results          google.Results
	Timeout, Elapsed time.Duration
}

func handleRequest(w http.ResponseWriter, req *http.Request) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	timeout, err := time.ParseDuration(req.FormValue("timeout"))
	if err != nil {
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}
	defer cancel()
	query := req.FormValue("q")
	if query == "" {
		http.Error(w, "no query", http.StatusBadRequest)
		return
	}
	userIP, err := userip.FromRequest(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ctx = userip.NewContext(ctx, userIP)
	// Run the Google search and print the results.
	start := time.Now()
	results, err := google.Search(ctx, query)
	elapsed := time.Since(start)
	if err := resultsTemplate.Execute(w, TemplateStruct{
		Results: results,
		Timeout: timeout,
		Elapsed: elapsed,
	}); err != nil {
		log.Print(err)
		return
	}
}

func main() {
	http.HandleFunc("/search", handleRequest)
	log.Fatal(http.ListenAndServe(":5055", nil))
}
