package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 6, 1, 0}
	fmt.Println(dominantIndex(nums))
}

func dominantIndex(nums []int) int {
	if (len(nums) == 1) {
		return 0
	}
	tempNums := make([]int,len(nums))
	copy(tempNums, nums)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if (nums[i] > nums[j]) {
				nums[i],nums[j] = nums[j],nums[i]
			}
		}
	}
	fmt.Println(nums)
	if (nums[0] >= nums[1]*2) {
		for i := 0; i < len(tempNums); i++ {
			if (tempNums[i] == nums[0]) {
				return i
			}
		}
	}
	return -1
}