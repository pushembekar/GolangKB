package main

import (
	"fmt"
)

func hello() {
	fmt.Println("hello ")
}

func world() {
	fmt.Println("world")
}

func main() {
	defer hello()
	defer fmt.Println("This is a test")
	defer world()
}
