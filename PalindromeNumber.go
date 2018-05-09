package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(10))
}

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	reslutStr := ""
	for i := len(str)-1; i >= 0; i-- {
		reslutStr += string(str[i])
	}
	if (reslutStr == str) {
		return true
	}else{
		return false
	}
}