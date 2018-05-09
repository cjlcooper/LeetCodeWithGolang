package main

import (
	"fmt"
)

func main() {
	J := "z"
	S := "aAAbbb"
	fmt.Println(numJewelsInStones(J, S))
}

func numJewelsInStones(J string, S string) int {
	nums := 0
	for i := 0; i < len(J); i++ {
		for j:= 0; j < len(S); j++ {
			if (string(J[i]) == string(S[j])) {
				nums++
			}
		}
	}

    return nums
}