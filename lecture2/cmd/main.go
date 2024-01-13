package main

import (
	"fmt"
	"os"

	"github.com/Praveen005/matt-holiday-go-lectures/tree/main/lecture2/hello"
)

func main(){
	//Args[0] is the name of program itself, what you type next becomes the next arguments

	// if len(os.Args) > 1 {
	// 	fmt.Println(hello.SayHello(os.Args[1:]))
	// }else{
	// 	fmt.Println(hello.SayHello("World"))
	// }

	fmt.Println(hello.SayHello(os.Args[1:]))
	
}