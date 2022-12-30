package main

import "fmt"

func main() {
	a := 100
	b := 200

	if a == 100 {
		if b == 200 {
			fmt.Printf("a 的值为 100 ， b 的值为 200\n")
		}
	}
	fmt.Printf("a 值为 : %d\n", a)
	fmt.Printf("b 值为 : %d\n", b)
}
