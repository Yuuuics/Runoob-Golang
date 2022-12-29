package main

import "fmt"

func main() {
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("c = %d\n", c)
	c = a - b
	fmt.Printf("c = %d\n", c)
	c = a * b
	fmt.Printf("c = %d\n", c)
	c = a / b
	fmt.Printf("c = %d\n", c)
	a++
	fmt.Printf("a + 1 = %d\n", a)
	b++
	fmt.Printf("b = %d\n", b)
}
