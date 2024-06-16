package main

import(
	"fmt"
)

func slice(){

	var f = make([]int, 5)
	fmt.Println(f)

	n := [] int {1,2,3}
	fmt.Println(n)
	p := make([]int, len(n))
	copy(p, n)

	fmt.Println(p)

	p= append(n[:0:0], n...)
	fmt.Println(p)

	p = append([]int(nil), n...)
	fmt.Println(p)

}