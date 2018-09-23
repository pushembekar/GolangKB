package main

import (
	"fmt"
)

func main() {
	a := 43

	fmt.Println(a)  //43
	fmt.Println(&a) // address of 43

	var b = &a // b has the address of a

	fmt.Println(b)  // tell me the address of a
	fmt.Println(*b) // tell me the value at address of a
}
