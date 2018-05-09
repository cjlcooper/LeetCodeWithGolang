package main

import (
	"fmt"
)

func main() {
	str := "hello"
	fmt.Println(reverseString(str))
}

func reverseString(s string) string {
    result := ""
    for i := len(s)-1; i >= 0; i-- {
    	result += string(s[i])
    }

    return result
}