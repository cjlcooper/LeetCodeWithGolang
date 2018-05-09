package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 3, 6}
	target := 6
	fmt.Println(towSum(nums, target))
}

func towSum(nums []int, target int) []int {
	var resultArr  = make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if (nums[i]+nums[j] == target && i!=j) {
				resultArr[0] = i
				resultArr[1] = j
			}
		}
	}
	if (resultArr[0] > resultArr[1]) {
		resultArr[0],resultArr[1] = resultArr[1],resultArr[0]
	}

	return resultArr
}