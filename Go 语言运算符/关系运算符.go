package main

import "fmt"

func main() {
	var a int = 21
	var b int = 10

	if a == b {
		fmt.Println("a = b\n")
	} else {
		fmt.Println("a != b")
	}

	if a < b {
		fmt.Println("a < b")
	} else {
		fmt.Println("a > b")
	}

	if a > b {
		fmt.Println("a > b")
	} else {
		fmt.Println("a < b")
	}

	// Change the value in a and b
	a = 5
	b = 20
	if a <= b {
		fmt.Println("a <= b")
	} else {
		fmt.Println("a >= b")
	}
}
