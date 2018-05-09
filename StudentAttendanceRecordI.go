package main

import (
	"fmt"
)

func main() {
	s := "PPALLL"
	fmt.Println(checkRecord(s))
}

func checkRecord(s string) bool {
	aCount := 0
	lCount := 0
	pCount := 0
	for i := 0; i < len(s); i++ {
		r := string(s[i])
		if (r == "A") {
			aCount++
		}else if (r == "L") {
			lCount++
		}else if (r == "P") {
			pCount++
		}
	}
	if (aCount > 1 || lCount>2) {
		return false
	}

	return true
} 