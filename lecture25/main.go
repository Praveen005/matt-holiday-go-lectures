// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"time"
// )

// type result struct {
// 	url     string
// 	err     error
// 	latency time.Duration
// }

// // old way
// // func get(url string, ch chan <- result) {
// // 	start := time.Now()

// // 	if resp, err := http.Get(url); err != nil{
// // 		ch <- result{url, err, 0}
// // 	}else{
// // 		t := time.Since(start).Round(time.Millisecond)
// // 		ch <- result{url, nil, t}
// // 		resp.Body.Close()
// // 	}
// // }

// // with context
// // what we did is just injected a 3sec timeout in our http request
// func get(ctx context.Context, url string, ch chan<- result) {
// 	start := time.Now()
// 	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
// 	if resp, err := http.DefaultClient.Do(req); err != nil {
// 		ch <- result{url, err, 0}
// 	} else {
// 		t := time.Since(start).Round(time.Millisecond)
// 		ch <- result{url, nil, t}
// 		resp.Body.Close()
// 	}
// }

// func main() {
// 	go func() {
// 		mux := http.NewServeMux()

// 		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 			fmt.Fprint(w, "Hello Praveen!")
// 		})

// 		_ = http.ListenAndServe(":8080", mux)
// 	}()

// 	results := make(chan result)
// 	list := []string{
// 		"https://amazon.com",
// 		"https://google.com",
// 		"https://nytimes.com",
// 		"https://wsj.com",
// 		"http://localhost:8080",
// 	}
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	for _, url := range list {
// 		go get(ctx, url, results)
// 	}

// 	for range list {
// 		r := <-results

// 		if r.err != nil {
// 			log.Printf("%-20s  %s\n", r.url, r.err)
// 		} else {
// 			log.Printf("%-20s  %s\n", r.url, r.latency)
// 		}
// 	}

// }

package main

import (
	"context"
	"log"
	"net/http"
	"runtime"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}



// with context
// what we did is just injected a 3sec timeout in our http request
func get(ctx context.Context, url string, ch chan<- result) {
	var r result
	start := time.Now()

	ticker := time.NewTicker(1 *time.Second).C

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if resp, err := http.DefaultClient.Do(req); err != nil {
		r = result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		r = result{url, nil, t}
		resp.Body.Close()
	}

	for{
		select{
		case ch <- r:
			return
		case <- ticker:
			log.Println("tick", r)
		}
	}
}

func first(ctx context.Context, urls []string)(*result, error) {
	results := make(chan result) // un-buffered

	// we want a cancel function
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, url := range urls{
		go get(ctx, url, results)
	}

	select{
	case r := <- results:
		return &r, nil // as soon as you get the first response, you will return and before returning, deferred cancellation will happen, resulting injust one response that we want
	case <- ctx.Done():
		return nil, ctx.Err()
	}
}

func main() {
	
	list := []string{
		"https://amazon.com",
		"https://google.com",
		"https://nytimes.com",
		"https://wsj.com",
	}

	// Get the first response and done!
	r, _ := first(context.Background(), list)

	if r.err != nil {
		log.Printf("%-20s  %s\n", r.url, r.err)
	} else {
		log.Printf("%-20s  %s\n", r.url, r.latency)
	}


	time.Sleep(9 *time.Second)
	log.Println("quit anyway...", runtime.NumGoroutine(), " still running")

}
