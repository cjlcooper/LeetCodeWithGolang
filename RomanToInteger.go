package main

import (
	"fmt"
)

func main() {
	//fmt.Println(romanToInt("VI"))
	//fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	dataDict := make(map[string]int)
	dataDict["I"] = 1
	dataDict["V"] = 5
	dataDict["X"] = 10
	dataDict["L"] = 50
	dataDict["C"] = 100
	dataDict["D"] = 500
	dataDict["M"] = 1000

	resultInt := 0
	tempResultInt := 0

	items := make([]int, len(s))
	//tempItems := make([]int,len(s))
	//add items slice
	for i := 0; i < len(s); i++ {
		items[i] = i
	}

	fmt.Println(items)
	//i := 1
	//items = append(items[:i], items[i+1:]...)

	for i := 0; i < len(s)-1; i++ {
		if (dataDict[string(s[i])]<dataDict[string(s[i+1])]) {
			r := dataDict[string(s[i+1])]-dataDict[string(s[i])]
			tempResultInt += r
		}
	}
	for i := 0; i < len(items); i++ {
		resultInt += dataDict[string(s[i])]
	}

 	return resultInt
}