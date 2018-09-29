package main

import (
	"fmt"
)

func main() {
	fmt.Println(greet("Jane ", "Doe"))
	fmt.Println(greeting(8, 9))
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}

func greeting(first, second int) string {
	return fmt.Sprint(first, second)
}
