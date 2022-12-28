// Sprintf 根据格式化参数生成格式化的字符串并返回该字符串。

package main

import "fmt"

func main() {
	var stockcode = 123
	var enddate = "2022-12-28"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)

	fmt.Println(target_url)
}
