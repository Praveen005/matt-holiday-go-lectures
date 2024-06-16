package main

import (
	"context"
	"fmt"
	"log"

	// "fmt"
	// "fmt"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

// old way
// func get(url string, ch chan <- result) {
// 	start := time.Now()

// 	if resp, err := http.Get(url); err != nil{
// 		ch <- result{url, err, 0}
// 	}else{
// 		t := time.Since(start).Round(time.Millisecond)
// 		ch <- result{url, nil, t}
// 		resp.Body.Close()
// 	}
// }


// with context
// what we did is just injected a 3sec timeout in our http request
func get(ctx context.Context, url string, ch chan <- result) {
	start := time.Now()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if resp, err := http.DefaultClient.Do(req); err != nil{
		ch <- result{url, err, 0}
	}else{
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t}
		resp.Body.Close()
	}
}

func main(){
	go func ()  {
		mux := http.NewServeMux()

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Hello Praveen!")
		})
	
		_ = http.ListenAndServe(":8080", mux)	
	}()
	
	results := make(chan result)
	list := []string{
		"https://amazon.com",
		"https://google.com",
		"https://nytimes.com",
		"https://wsj.com",
		"http://localhost:8080",
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3 *time.Second)
	defer cancel()

	for _, url := range list{
		go get(ctx, url, results)
	}

	

	for range list{
		r := <- results

		if r.err != nil{
			log.Printf("%-20s  %s\n", r.url, r.err)
		}else{
			log.Printf("%-20s  %s\n", r.url, r.latency)
		}
	}
	
}