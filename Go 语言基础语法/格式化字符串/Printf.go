package main

import "fmt"

func main() {
	var stockcode = 123
	var enddate = "2022-12-28"
	var url = "Code=%d&endDate=%s"

	fmt.Printf(url, stockcode, enddate)
}
