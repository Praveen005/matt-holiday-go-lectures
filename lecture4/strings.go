package main

import(
	"bufio"
	"fmt"
	"strings"
	"os"
)

func stringOp(){
	if len(os.Args) <3 {
		fmt.Fprintln(os.Stderr, "Not enough args")
		return
	}

	old, new := os.Args[1], os.Args[2]

	//creating a scanner: it is just a buffered io tool around the standarad input and by default, it always split the input by lines

	scan := bufio.n
}