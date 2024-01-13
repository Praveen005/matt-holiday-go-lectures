package main

import(
	"fmt"
	"os"
)

func avg(){
	var sum float64
	var n int
	for{
		var val float64

		if _, err := fmt.Fscanln(os.Stdin, &val); err != nil{
			break
		}

		sum += val
		n++ 
	}

	if n== 0{
		fmt.Fprintln(os.Stderr, "no values")
		os.Exit(-1)  //exit and panic are best to be used only in main()
	}

	fmt.Printf("The average of %d numbers is: %f\n", n, sum/float64(n)) //need to cast n first, since sum is a float
}