package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d - %b - %x\n", 42, 42, 42)
	fmt.Printf("%d - %b - %#x\n", 42, 42, 42)
	fmt.Printf("%d\t%b\t%o\t%#o\t%x\t%#x\n", 42, 42, 42, 42, 42, 42)
}
