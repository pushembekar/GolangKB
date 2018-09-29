package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("This ran")
	}

	b := true

	if food := "Chocolate"; b {
		fmt.Println(food)
	}
}
