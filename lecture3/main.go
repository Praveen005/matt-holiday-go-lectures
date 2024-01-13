package main

import "fmt"

//Following declaration can be made anywhere

var a int

var(
	b = 2 		//Here go will infer the type automatically, int/int64
	f= 4.3 		//float64
)

func main(){

	//This kinda declaration can only be made inside a function
	//Short declaration operator :=
	c := 54

	fmt.Printf("c: %9T %v\n", c, c)
	fmt.Printf("f: %8T %v\n", f, f)

	//[1] means use whatever the parameter 1 was, so you don't need to do, fmt.Printf("b: %8T %[1]v\n", b, b)
	fmt.Printf("b: %8T %[1]v\n", b)
	fmt.Printf("c: %9T %[1]v\n", c, c)
	fmt.Printf("f: %8T %[1]v\n", f, f)

	//can't assign one type to another, need to do type conversion

	a= int(f)
	fmt.Println(a)

	avg(); //press ctrl+c(ctrl + d in linux) to stop reading the input and calculate the average

}