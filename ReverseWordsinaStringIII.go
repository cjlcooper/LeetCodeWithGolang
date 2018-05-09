package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Let's take LeetCode contest"
	fmt.Println(reverseWords(s))
}

func reverseWords(s string) string {
	resultStr := ""
	splitArr := strings.Split(s, " ")
	for i := 0; i < len(splitArr); i++ {
		singleStr := splitArr[i]
		if (i != 0) {
			resultStr += " "
		}
		for j := len(singleStr)-1; j >= 0; j-- {
			resultStr += string(singleStr[j])
		}
	}
	return resultStr
}