package main

import "fmt"

func main() {
	a := true
	b := false

	if a && b {
		fmt.Println("true")
	}
	if a || b {
		fmt.Println("true")
	}

	// change the value in a and b
	a = false
	b = true
	if a && b {
		fmt.Println("ture")
	} else {
		fmt.Println("false")
	}
	if !(a && b) {
		fmt.Println("true")
	}
}
