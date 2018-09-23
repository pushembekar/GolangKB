package main

import (
	"fmt"
)

func main() {
	a := 43
	fmt.Println("a - ", a)
	fmt.Println("a's memory is: ", &a)
	fmt.Printf("a's memory is: %d", &a)
}
