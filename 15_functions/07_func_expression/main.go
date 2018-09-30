package main

import (
	"fmt"
)

func main() {
	greeting := func() {
		fmt.Println("Hello world!")
	}

	greeting()
	fmt.Printf("%T\n", greeting)

	greet := makeGreeter()
	fmt.Println(greet())
}

func makeGreeter() func() string {
	return func() string {
		return "Hello World!"
	}
}
