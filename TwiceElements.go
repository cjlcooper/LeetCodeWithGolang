package main

import (
	"fmt"
)

func main() {
	nums := []int{4,3,2,7,8,2,3,1}
	fmt.Println(findDuplicates(nums))
}

func findDuplicates(nums []int) []int {
	if (len(nums) <= 1) {
		return nil
	}
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
	return nil
}