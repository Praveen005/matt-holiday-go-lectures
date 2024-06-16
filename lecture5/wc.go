package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)


func count(){
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)

	// analyze my scanner with a different spiliting function
	// It will give us one word at a time
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		words[scan.Text()]++
	}

	if err := scan.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		return
	}

	fmt.Println(len(words), "unique words")

	type kv struct{
		key		string
		val		int
	}

	var ss []kv

	for k, v := range words{
		ss = append(ss, kv{k,v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].val > ss[j].val  // want decreasing order
	})


	for _, s := range ss[:3] {
		fmt.Println(s.key, " appears ", s.val, " times.")
	}
}

