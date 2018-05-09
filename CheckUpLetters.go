package main

import(
	"fmt"
	"strings"
)

func  main() {
	fmt.Println(detectCapitalUse("Google"))
}

func detectCapitalUse(word string) bool {
	if (len(word) == 0) {
		return true
	}

	//upper
	afterUpperStr := strings.ToUpper(word)
	if (afterUpperStr == word) {
		return true
	}

	//lower
	afterLowerStr := strings.ToLower(word)
	if (afterLowerStr == word) {
		return true
	}

	//first is upper
	firstItemStr := string([]byte(word)[:1])
	firstItemUpperStr := strings.ToUpper(firstItemStr)
	otherItemsStr := word[1:len(word)]
	otherItemsLowerStr := strings.ToLower(otherItemsStr)
	if (firstItemUpperStr == firstItemStr && otherItemsLowerStr == otherItemsStr) {
		return true
	}


	return false
}