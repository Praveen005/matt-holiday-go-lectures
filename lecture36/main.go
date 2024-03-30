package main

import "net/http"

const url = "http://jsonplaceholder.typicode.com"

type todo struct {
	UserID    int    `json:"userID"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var mark = map[bool]string{
	false: " ",
	true:  "x",
}

func handler(w http.ResponseWriter, r *http.Request) {

	req, _ := http.NewRequest("GET", url+"/todos/"+r.URL.Path[1:], nil)
	tr := &http.Transport{}
	cli := &http.Client{Transport: tr}
	resp, err := cli.Do(req)
}
