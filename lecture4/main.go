package main

import(
	"fmt"
)

func main(){
	s := "élite"
	fmt.Printf("s: %8T %[1]v\n", s)
	fmt.Printf("s: %8T %[1]v\n", []rune(s))  //we have 5 characters in 's', so we have five values in the slice of runes
	fmt.Printf("s: %8T %[1]v\n", []byte(s))

}

/*
Lets's understand the output:
s:   string élite
s:  []int32 [233 108 105 116 101]
s:  []uint8 [195 169 108 105 116 101]

- first value of the slice of runes is 233, remember ASCII characters ranges from 0 to 127, but 'é' is not an ascii character
  But rather a french 'e' and we know, rune is a 32-bit or 4 byte integer, it prints 233

- See, in slice of bytes, which is just 1 byte or 8 bit integer, which can store from 0 to 127 and hence can't store 233, so this slice of byte contains 6 values and not just 5 for 5 characters in string.
*/