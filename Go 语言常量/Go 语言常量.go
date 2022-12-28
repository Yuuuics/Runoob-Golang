package main

import "fmt"

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	const a, b, c = 1, false, "str" // 多重赋值
	var area int

	area = LENGTH * WIDTH
	fmt.Printf("面积为: %d", area)
	println()
	println(a, b, c)
}
