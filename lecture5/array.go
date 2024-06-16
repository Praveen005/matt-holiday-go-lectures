package main

import(
	"fmt"
)

func arr(){
	var a [4]int
	//Following is the way to declare an array, . . . means, its size will be equal to the no. of elements in it
	var c = [...] int {0, 0, 0}

	m := [...] int {1, 2, 3, 4}

	// c= m  // TYPE MISMATCH, their sizes are different, so can't do it
	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(m)


}