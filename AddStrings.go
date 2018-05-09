package main

import (
	"fmt"
	"strconv"
)

func main() {
	num1 := "9999"
	num2 := "0"
	fmt.Println(addStrings(num1, num2))
}

func addStrings(num1 string, num2 string) string {
	int1, err := strconv.ParseInt(num1, 10, 64)
	int2, err := strconv.ParseInt(num2, 10, 64)
	if (nil != err) {
		return ""
	}
	resultInt := int1 + int2
	resultStr := strconv.FormatInt(resultInt, 10)

	return resultStr
}