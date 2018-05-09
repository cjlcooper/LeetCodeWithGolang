package main

import (
	"fmt"
)

func main() {
	nums := []int{2,0,2,1,1,0}
	sortColors(nums)
}

func sortColors(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if (nums[i] < nums[j]) {
				nums[i],nums[j] = nums[j],nums[i]
			}
		}
	}
	fmt.Println(nums)
}